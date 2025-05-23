package lxd

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/canonical/lxd/shared/api"
)

// Certificate handling functions

// GetCertificateFingerprints returns a list of certificate fingerprints.
func (r *ProtocolLXD) GetCertificateFingerprints() ([]string, error) {
	// Fetch the raw URL values.
	urls := []string{}
	baseURL := "/certificates"
	_, err := r.queryStruct(http.MethodGet, baseURL, nil, "", &urls)
	if err != nil {
		return nil, err
	}

	// Parse it.
	return urlsToResourceNames(baseURL, urls...)
}

// GetCertificates returns a list of certificates.
func (r *ProtocolLXD) GetCertificates() ([]api.Certificate, error) {
	certificates := []api.Certificate{}

	// Fetch the raw value
	_, err := r.queryStruct(http.MethodGet, "/certificates?recursion=1", nil, "", &certificates)
	if err != nil {
		return nil, err
	}

	return certificates, nil
}

// GetCertificate returns the certificate entry for the provided fingerprint.
func (r *ProtocolLXD) GetCertificate(fingerprint string) (*api.Certificate, string, error) {
	certificate := api.Certificate{}

	// Fetch the raw value
	etag, err := r.queryStruct(http.MethodGet, "/certificates/"+url.PathEscape(fingerprint), nil, "", &certificate)
	if err != nil {
		return nil, "", err
	}

	return &certificate, etag, nil
}

// CreateCertificate adds a new certificate to the LXD trust store.
func (r *ProtocolLXD) CreateCertificate(certificate api.CertificatesPost) error {
	// Send the request
	_, _, err := r.query(http.MethodPost, "/certificates", certificate, "")
	if err != nil {
		return err
	}

	return nil
}

// UpdateCertificate updates the certificate definition.
func (r *ProtocolLXD) UpdateCertificate(fingerprint string, certificate api.CertificatePut, ETag string) error {
	err := r.CheckExtension("certificate_update")
	if err != nil {
		return err
	}

	// Send the request
	_, _, err = r.query(http.MethodPut, "/certificates/"+url.PathEscape(fingerprint), certificate, ETag)
	if err != nil {
		return err
	}

	return nil
}

// DeleteCertificate removes a certificate from the LXD trust store.
func (r *ProtocolLXD) DeleteCertificate(fingerprint string) error {
	// Send the request
	_, _, err := r.query(http.MethodDelete, "/certificates/"+url.PathEscape(fingerprint), nil, "")
	if err != nil {
		return err
	}

	return nil
}

// CreateCertificateToken requests a certificate add token.
func (r *ProtocolLXD) CreateCertificateToken(certificate api.CertificatesPost) (Operation, error) {
	err := r.CheckExtension("certificate_token")
	if err != nil {
		return nil, err
	}

	if !certificate.Token {
		return nil, errors.New("Token needs to be true if requesting a token")
	}

	// Send the request
	op, _, err := r.queryOperation(http.MethodPost, "/certificates", certificate, "", true)
	if err != nil {
		return nil, err
	}

	return op, nil
}
