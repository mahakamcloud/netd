package config

import (
	"errors"
	"fmt"
	"net"

	"github.com/mahakamcloud/mahakam/pkg/utils"
	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

// Config represents mahakam configuration
type Config struct {
	MahakamServerConfig MahakamServerConfig  `yaml:"server"`
	LogLevel            string               `yaml:"log_level"`
	KVStoreConfig       StorageBackendConfig `yaml:"storage_backend"`
	NetworkConfig       NetworkConfig        `yaml:"network"`
	KubernetesConfig    KubernetesConfig     `yaml:"kubernetes"`
	GateConfig          GateConfig           `yaml:"gate"`
	TerraformConfig     TerraformConfig      `yaml:"terraform"`
	HostsConfig         []Host               `yaml:"hosts"`
}

// LoadConfig loads a configuration file
func LoadConfig(configFilePath string) (*Config, error) {
	var config *Config
	if configFilePath == "" {
		return config, fmt.Errorf("must provide non-empty configuration file path")
	}

	bytes, err := utils.ReadFile(configFilePath)
	if err != nil {
		return config, err
	}

	if err = yaml.Unmarshal(bytes, &config); err != nil {
		return config, fmt.Errorf("error unmarshaling configuration file: %s", err)
	}

	if err = config.Validate(); err != nil {
		return config, fmt.Errorf("error validating configuration file: %s", err)
	}

	return config, nil
}

// Validate validates mahakam configuration
func (c *Config) Validate() error {
	if c.LogLevel == "" {
		return fmt.Errorf("must provide non-empty log level")
	}

	if err := c.MahakamServerConfig.Validate(); err != nil {
		return fmt.Errorf("error validating mahakam server configuration: %s", err)
	}

	if err := c.KVStoreConfig.Validate(); err != nil {
		return fmt.Errorf("error validating KV store configuration: %s", err)
	}

	if err := c.NetworkConfig.Validate(); err != nil {
		return fmt.Errorf("error validating network configuration: %s", err)
	}

	if err := c.TerraformConfig.Validate(); err != nil {
		return fmt.Errorf("error validating terraform configuration: %s", err)
	}

	if err := c.GateConfig.Validate(); err != nil {
		return fmt.Errorf("error validating gate configuration: %s", err)
	}

	if len(c.HostsConfig) < 1 {
		return fmt.Errorf("error validating hosts configuration: %s", errors.New("empty hosts list"))
	}

	for _, host := range c.HostsConfig {
		if err := host.Validate(); err != nil {
			return fmt.Errorf("error validating hosts configuration: %s", err)
		}
	}

	return nil
}

// MahakamServerConfig stores mahakam server config
type MahakamServerConfig struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

// Validate checks mahakam server configuration
func (m *MahakamServerConfig) Validate() error {
	if m.Port == 0 {
		return fmt.Errorf("must provide non-empty port")
	}

	return nil
}

// StorageBackendConfig stores metadata for storage backend that we use
type StorageBackendConfig struct {
	BackendType string `yaml:"backend_type"`
	Address     string `yaml:"address"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Bucket      string `yaml:"bucket"`
}

// Validate validates storage backend configuration
func (sbc *StorageBackendConfig) Validate() error {
	if sbc.BackendType == "" {
		return fmt.Errorf("must provide non-empty Backend type")
	}

	if sbc.Address == "" {
		return fmt.Errorf("must provide non-empty storage backend address")
	}

	return nil
}

// CheckStorageBackendConnection defines connect check on StorageBackendConfig
type CheckStorageBackendConnection struct {
	StorageBackendConfig
	log         logrus.FieldLogger
	pingChecker utils.PingChecker
}

// NewCheckStorageBackendConnection return new CheckStorageBackendConnection
func NewCheckStorageBackendConnection(s StorageBackendConfig, log logrus.FieldLogger, pingChecker utils.PingChecker) *CheckStorageBackendConnection {
	return &CheckStorageBackendConnection{
		StorageBackendConfig: s,
		log:                  log,
		pingChecker:          pingChecker,
	}
}

// ValidateAvailability validates storage backend reachability
func (c *CheckStorageBackendConnection) ValidateAvailability() error {
	pingCheck := utils.NewPingCheck()

	backendReady := pingCheck.PortPingNWithDelay(c.Address, StorageBackendPingTimeout, c.log, StorageBackendPingRetry, StorageBackendPingDelay)

	// Storage backend still not ready after max retry
	if !backendReady {
		return fmt.Errorf("timeout waiting for storage backend to be up '%v'", c.Address)
	}

	return nil
}

// NetworkConfig stores metadata for network that mahakam will configure
type NetworkConfig struct {
	// CIDR is datacenter network CIDR that Mahakam will use to provision cluster network from it
	CIDR string `yaml:"cidr"`
	// ClusterNetmask is subnet length that cluster network will be provisioned as
	ClusterNetmask int `yaml:"cluster_netmask"`
	// DatacenterGatewayCIDR is gateway public IP in the datacenter that can reach Internet
	DatacenterGatewayCIDR string `yaml:"datacenter_gateway_cidr"`
	// DatacenterNameserver is nameserver in datacenter that can solve domains in Internet
	DatacenterNameserver string `yaml:"datacenter_nameserver"`
	// Domain is domain for the network
	Domain string `yaml:"domain"`
}

// Validate validates storage backend configuration
func (nc *NetworkConfig) Validate() error {
	if nc.CIDR == "" {
		return fmt.Errorf("must provide non-empty network CIDR")
	}

	if nc.ClusterNetmask == 0 {
		return fmt.Errorf("must provide non-empty cluster netmask")
	}

	if _, _, err := net.ParseCIDR(nc.CIDR); err != nil {
		return fmt.Errorf("must provide valid CIDR format for cluster network")
	}

	if nc.ClusterNetmask > 32 || nc.ClusterNetmask < 1 {
		return fmt.Errorf("must provide valid cluster netmask between 0 and 32")
	}

	if nc.DatacenterGatewayCIDR == "" {
		return fmt.Errorf("must provide non-empty datacenter gateway CIDR")
	}

	if _, _, err := net.ParseCIDR(nc.DatacenterGatewayCIDR); err != nil {
		return fmt.Errorf("must provide valid CIDR format for datacenter gateway")
	}

	if nc.DatacenterNameserver == "" {
		return fmt.Errorf("must provide non-empty datacenter nameserver")
	}

	if validIP := net.ParseIP(nc.DatacenterNameserver); validIP == nil {
		return fmt.Errorf("must provide valid IP format for datacenter nameserver")
	}

	if nc.Domain == "" {
		return fmt.Errorf("must provide non-empty domain")
	}

	return nil
}

// KubernetesConfig stores metadata for kubernetes cluster that mahakam will configure
type KubernetesConfig struct {
	// PodNetworkCidr is pod network that CNI will provision inside Kubernetes cluster
	PodNetworkCidr string `yaml:"pod_network_cidr"`
	// KubeadmToken is token secret used for workers to join Kubernetes cluster
	KubeadmToken string `yaml:"kubeadm_token"`
	// SSHPublicKey is public key that will be authorized to access Kubernetes nodes from its private key pair
	SSHPublicKey string `yaml:"ssh_public_key"`
}

// Validate validates storage backend configuration
func (kc *KubernetesConfig) Validate() error {
	if kc.PodNetworkCidr == "" {
		return fmt.Errorf("must provide non-empty pod network CIDR")
	}

	if kc.KubeadmToken == "" {
		return fmt.Errorf("must provide non-empty kubeadm token")
	}

	if kc.SSHPublicKey == "" {
		return fmt.Errorf("must provide non-empty SSH public key")
	}

	return nil
}

type GateConfig struct {
	GateNSSApiKey string `yaml:"gate_nss_api_key"`
}

func (gc *GateConfig) Validate() error {
	if gc.GateNSSApiKey == "" {
		return fmt.Errorf("must provide non-empty Gate NSS API key")
	}
	return nil
}

type TerraformConfig struct {
	LibvirtModulePath       string `yaml:"libvirt_module_path"`
	PublicLibvirtModulePath string `yaml:"public_libvirt_module_path"`
	ImageSourcePath         string `yaml:"image_source_path"`
}

func (tc *TerraformConfig) Validate() error {
	if tc.LibvirtModulePath == "" {
		return fmt.Errorf("must provide non-empty libvirt module path")
	}
	if tc.PublicLibvirtModulePath == "" {
		return fmt.Errorf("must provide non-empty public libvirt module path")
	}
	if tc.ImageSourcePath == "" {
		return fmt.Errorf("must provide non-empty image source path")
	}
	return nil
}

type Host struct {
	Name      string `yaml:"name"`
	IPAddress string `yaml:"ip_address"`
}

func (h *Host) Validate() error {
	if h.Name == "" {
		return fmt.Errorf("must provide non-empty host name for host: [%s]", h)
	}

	if validIP := net.ParseIP(h.IPAddress); validIP == nil {
		return fmt.Errorf("must provide valid IP format for host: [%s]", h)
	}
	return nil
}
