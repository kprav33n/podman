// Code generated by go generate; DO NOT EDIT.
package play

import (
	"net"
	"net/url"

	"github.com/containers/podman/v3/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *KubeOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *KubeOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithAuthfile set field Authfile to given value
func (o *KubeOptions) WithAuthfile(value string) *KubeOptions {
	o.Authfile = &value
	return o
}

// GetAuthfile returns value of field Authfile
func (o *KubeOptions) GetAuthfile() string {
	if o.Authfile == nil {
		var z string
		return z
	}
	return *o.Authfile
}

// WithCertDir set field CertDir to given value
func (o *KubeOptions) WithCertDir(value string) *KubeOptions {
	o.CertDir = &value
	return o
}

// GetCertDir returns value of field CertDir
func (o *KubeOptions) GetCertDir() string {
	if o.CertDir == nil {
		var z string
		return z
	}
	return *o.CertDir
}

// WithUsername set field Username to given value
func (o *KubeOptions) WithUsername(value string) *KubeOptions {
	o.Username = &value
	return o
}

// GetUsername returns value of field Username
func (o *KubeOptions) GetUsername() string {
	if o.Username == nil {
		var z string
		return z
	}
	return *o.Username
}

// WithPassword set field Password to given value
func (o *KubeOptions) WithPassword(value string) *KubeOptions {
	o.Password = &value
	return o
}

// GetPassword returns value of field Password
func (o *KubeOptions) GetPassword() string {
	if o.Password == nil {
		var z string
		return z
	}
	return *o.Password
}

// WithNetwork set field Network to given value
func (o *KubeOptions) WithNetwork(value string) *KubeOptions {
	o.Network = &value
	return o
}

// GetNetwork returns value of field Network
func (o *KubeOptions) GetNetwork() string {
	if o.Network == nil {
		var z string
		return z
	}
	return *o.Network
}

// WithNoHosts set field NoHosts to given value
func (o *KubeOptions) WithNoHosts(value bool) *KubeOptions {
	o.NoHosts = &value
	return o
}

// GetNoHosts returns value of field NoHosts
func (o *KubeOptions) GetNoHosts() bool {
	if o.NoHosts == nil {
		var z bool
		return z
	}
	return *o.NoHosts
}

// WithQuiet set field Quiet to given value
func (o *KubeOptions) WithQuiet(value bool) *KubeOptions {
	o.Quiet = &value
	return o
}

// GetQuiet returns value of field Quiet
func (o *KubeOptions) GetQuiet() bool {
	if o.Quiet == nil {
		var z bool
		return z
	}
	return *o.Quiet
}

// WithSignaturePolicy set field SignaturePolicy to given value
func (o *KubeOptions) WithSignaturePolicy(value string) *KubeOptions {
	o.SignaturePolicy = &value
	return o
}

// GetSignaturePolicy returns value of field SignaturePolicy
func (o *KubeOptions) GetSignaturePolicy() string {
	if o.SignaturePolicy == nil {
		var z string
		return z
	}
	return *o.SignaturePolicy
}

// WithSkipTLSVerify set field SkipTLSVerify to given value
func (o *KubeOptions) WithSkipTLSVerify(value bool) *KubeOptions {
	o.SkipTLSVerify = &value
	return o
}

// GetSkipTLSVerify returns value of field SkipTLSVerify
func (o *KubeOptions) GetSkipTLSVerify() bool {
	if o.SkipTLSVerify == nil {
		var z bool
		return z
	}
	return *o.SkipTLSVerify
}

// WithSeccompProfileRoot set field SeccompProfileRoot to given value
func (o *KubeOptions) WithSeccompProfileRoot(value string) *KubeOptions {
	o.SeccompProfileRoot = &value
	return o
}

// GetSeccompProfileRoot returns value of field SeccompProfileRoot
func (o *KubeOptions) GetSeccompProfileRoot() string {
	if o.SeccompProfileRoot == nil {
		var z string
		return z
	}
	return *o.SeccompProfileRoot
}

// WithStaticIPs set field StaticIPs to given value
func (o *KubeOptions) WithStaticIPs(value []net.IP) *KubeOptions {
	o.StaticIPs = &value
	return o
}

// GetStaticIPs returns value of field StaticIPs
func (o *KubeOptions) GetStaticIPs() []net.IP {
	if o.StaticIPs == nil {
		var z []net.IP
		return z
	}
	return *o.StaticIPs
}

// WithStaticMACs set field StaticMACs to given value
func (o *KubeOptions) WithStaticMACs(value []net.HardwareAddr) *KubeOptions {
	o.StaticMACs = &value
	return o
}

// GetStaticMACs returns value of field StaticMACs
func (o *KubeOptions) GetStaticMACs() []net.HardwareAddr {
	if o.StaticMACs == nil {
		var z []net.HardwareAddr
		return z
	}
	return *o.StaticMACs
}

// WithConfigMaps set field ConfigMaps to given value
func (o *KubeOptions) WithConfigMaps(value []string) *KubeOptions {
	o.ConfigMaps = &value
	return o
}

// GetConfigMaps returns value of field ConfigMaps
func (o *KubeOptions) GetConfigMaps() []string {
	if o.ConfigMaps == nil {
		var z []string
		return z
	}
	return *o.ConfigMaps
}

// WithLogDriver set field LogDriver to given value
func (o *KubeOptions) WithLogDriver(value string) *KubeOptions {
	o.LogDriver = &value
	return o
}

// GetLogDriver returns value of field LogDriver
func (o *KubeOptions) GetLogDriver() string {
	if o.LogDriver == nil {
		var z string
		return z
	}
	return *o.LogDriver
}

// WithLogOptions set field LogOptions to given value
func (o *KubeOptions) WithLogOptions(value []string) *KubeOptions {
	o.LogOptions = &value
	return o
}

// GetLogOptions returns value of field LogOptions
func (o *KubeOptions) GetLogOptions() []string {
	if o.LogOptions == nil {
		var z []string
		return z
	}
	return *o.LogOptions
}

// WithStart set field Start to given value
func (o *KubeOptions) WithStart(value bool) *KubeOptions {
	o.Start = &value
	return o
}

// GetStart returns value of field Start
func (o *KubeOptions) GetStart() bool {
	if o.Start == nil {
		var z bool
		return z
	}
	return *o.Start
}
