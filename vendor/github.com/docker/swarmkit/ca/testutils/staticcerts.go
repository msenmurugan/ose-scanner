package testutils

var (
	// NotYetValidCert is an ECDSA CA certificate that becomes valid in 2117, and expires in 2316
	NotYetValidCert = []byte(`
-----BEGIN CERTIFICATE-----
MIIBajCCARCgAwIBAgIUWYyg+FvrTJ/wtJd4pZF/GfO5uC0wCgYIKoZIzj0EAwIw
ETEPMA0GA1UEAxMGcm9vdENOMCIYDzIxMTcwMTAyMTgxODUyWhgPMjMxNjExMTUx
ODE4NTJaMBExDzANBgNVBAMTBnJvb3RDTjBZMBMGByqGSM49AgEGCCqGSM49AwEH
A0IABDC0qWmbfAkZH01xUVjwwR+2ovotU1iVIUD2fOFm93WUfg31cyga9dPDsg7R
GXJlRBnU9A48TWZMzIcqaa9ZpwyjQjBAMA4GA1UdDwEB/wQEAwIBBjAPBgNVHRMB
Af8EBTADAQH/MB0GA1UdDgQWBBS17zzXe1Q2tBZGw8xGL0spE88yQTAKBggqhkjO
PQQDAgNIADBFAiEAvnTTPh/jgnXIyLmbfROftfY2zCk0C0XLfLVnSj5MDZwCIDdP
tPG9bWx1C0I55UiWXKGQf3nUU68nQkk9JxVyjBma
-----END CERTIFICATE-----
`)
	// NotYetValidKey is the key corresponding to the NotYetValidCert
	NotYetValidKey = []byte(`
-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIOPDjFG/meAtl1a/mXP66Y17O7TFCop9JXo5Il90qYLNoAoGCCqGSM49
AwEHoUQDQgAEMLSpaZt8CRkfTXFRWPDBH7ai+i1TWJUhQPZ84Wb3dZR+DfVzKBr1
08OyDtEZcmVEGdT0DjxNZkzMhyppr1mnDA==
-----END EC PRIVATE KEY-----
`)

	// ExpiredCert is an ECDSA CA certificate that expired in 2007 (1967-2007)
	ExpiredCert = []byte(`
-----BEGIN CERTIFICATE-----
MIIBZzCCAQygAwIBAgIUNwwbocQMXzakEpwZoGkk7yOleRgwCgYIKoZIzj0EAwIw
ETEPMA0GA1UEAxMGcm9vdENOMB4XDTY3MDIyNDIzMDc0MFoXDTA3MDIyNDIzMDc0
MFowETEPMA0GA1UEAxMGcm9vdENOMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE
MLSpaZt8CRkfTXFRWPDBH7ai+i1TWJUhQPZ84Wb3dZR+DfVzKBr108OyDtEZcmVE
GdT0DjxNZkzMhyppr1mnDKNCMEAwDgYDVR0PAQH/BAQDAgEGMA8GA1UdEwEB/wQF
MAMBAf8wHQYDVR0OBBYEFLXvPNd7VDa0FkbDzEYvSykTzzJBMAoGCCqGSM49BAMC
A0kAMEYCIQCx5Lhl4b3YsjQuqHT/+vL5rnc0GV/OwJ8l2GFS2IB7EgIhAKrHZrcr
5+MmM1YUiykjweok2j5rj0/+9sR7waa69dkW
-----END CERTIFICATE-----
`)
	// ExpiredKey is the key corresponding to the ExpiredCert
	ExpiredKey = []byte(`
-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIOPDjFG/meAtl1a/mXP66Y17O7TFCop9JXo5Il90qYLNoAoGCCqGSM49
AwEHoUQDQgAEMLSpaZt8CRkfTXFRWPDBH7ai+i1TWJUhQPZ84Wb3dZR+DfVzKBr1
08OyDtEZcmVEGdT0DjxNZkzMhyppr1mnDA==
-----END EC PRIVATE KEY-----
`)

	// RSA2048SHA256Cert is an RSA CA cert with a 2048-bit key, SHA256 signature algorithm, that is currently valid and expires in 2117.
	// This should be valid because the key length is at least 2048 and the signature algorithm is SHA256.
	RSA2048SHA256Cert = []byte(`
-----BEGIN CERTIFICATE-----
MIIDjDCCAnSgAwIBAgIJAI5MpW7XttrnMA0GCSqGSIb3DQEBCwUAMGExCzAJBgNV
BAYTAlVTMQswCQYDVQQIEwJDQTEWMBQGA1UEBxMNU2FuIEZyYW5jaXNjbzEPMA0G
A1UEChMGRG9ja2VyMRwwGgYDVQQDExNTd2FybWtpdCBDQSBUZXN0aW5nMCAXDTE3
MDEyNzAwMzM1N1oYDzIxMTcwMTAzMDAzMzU3WjBhMQswCQYDVQQGEwJVUzELMAkG
A1UECBMCQ0ExFjAUBgNVBAcTDVNhbiBGcmFuY2lzY28xDzANBgNVBAoTBkRvY2tl
cjEcMBoGA1UEAxMTU3dhcm1raXQgQ0EgVGVzdGluZzCCASIwDQYJKoZIhvcNAQEB
BQADggEPADCCAQoCggEBAOj20YFx3Lo3xxshGwCirWixp3Wxa+k3Fpa3o1fZ3+jl
1V0op4swrf9EckrrDNHSkDeWRVjSuYAZ/t8KT/B/JuP8rL8PWkQD0BUQAuArxIsG
JuhfwuNhM6mhSEEMIrb6g1XLQ37rW5a9FTIbY+QJgYsPgWjFRgY5cT+ZXrgacmg6
cVWF75wSjW5DzZavGVfHPDebl0dXqeUHXvksZZ/pfzsTyqlgVp3Br5PKON6UqHNT
zI8MWEeTT+jpFTSR4Qt/Gdp5PbzTxfun38oOgT8WB3xJ1XvrRsxROPluBa1y7cVm
UcriTPzUtAhxb7MVGaTVwQ1zX1Wd+t0mYQVW8zRMK6ECAwEAAaNFMEMwEgYDVR0T
AQH/BAgwBgEB/wIBATAOBgNVHQ8BAf8EBAMCAUYwHQYDVR0OBBYEFLzBQbsbg8is
pyWorw6eP2ftJETsMA0GCSqGSIb3DQEBCwUAA4IBAQAPj5P1v1fqxUSs/uswfNZ2
APb7h1bccP41bEmgX45m0g7S4fLoFZb501IzgF6fsmJibhOJ/mKrPi5VM1RFpMfM
mL5zpdEXsopIfn9J4liXGXM1gFH6s4GeEn6cIwT7Sfzo1VPS0qbe9KJqPCLFySev
DivyL8Yv/NbTPF1wTrtoAhQeADSMxdctTutLMKE4CbJWhSPpvnojL94Jxj5TkUKR
fpg1gDGYtAcxpE+qZBI+YCh0r9ae/Wtg3lzw+I7/usmfO2Pm56Hb/O7ulRuLEOFu
XL2VZMKBpOTyDpe3YXMcvp3HT4qO5PmNs1b/N3Q8GwYRwfg6DZX2fPHT9vJGEdyq
-----END CERTIFICATE-----
`)
	// RSA2048SHA1Cert is an RSA CA cert with a 2048-bit key, SHA1 signature algorithm, that is currently valid and expires in 2117.
	// This should be not valid because the signature algorithm is SHA1.
	RSA2048SHA1Cert = []byte(`
-----BEGIN CERTIFICATE-----
MIIDjDCCAnSgAwIBAgIJAI6dSku42a9hMA0GCSqGSIb3DQEBBQUAMGExCzAJBgNV
BAYTAlVTMQswCQYDVQQIEwJDQTEWMBQGA1UEBxMNU2FuIEZyYW5jaXNjbzEPMA0G
A1UEChMGRG9ja2VyMRwwGgYDVQQDExNTd2FybWtpdCBDQSBUZXN0aW5nMCAXDTE3
MDEyNzAwMzQzNloYDzIxMTcwMTAzMDAzNDM2WjBhMQswCQYDVQQGEwJVUzELMAkG
A1UECBMCQ0ExFjAUBgNVBAcTDVNhbiBGcmFuY2lzY28xDzANBgNVBAoTBkRvY2tl
cjEcMBoGA1UEAxMTU3dhcm1raXQgQ0EgVGVzdGluZzCCASIwDQYJKoZIhvcNAQEB
BQADggEPADCCAQoCggEBAOj20YFx3Lo3xxshGwCirWixp3Wxa+k3Fpa3o1fZ3+jl
1V0op4swrf9EckrrDNHSkDeWRVjSuYAZ/t8KT/B/JuP8rL8PWkQD0BUQAuArxIsG
JuhfwuNhM6mhSEEMIrb6g1XLQ37rW5a9FTIbY+QJgYsPgWjFRgY5cT+ZXrgacmg6
cVWF75wSjW5DzZavGVfHPDebl0dXqeUHXvksZZ/pfzsTyqlgVp3Br5PKON6UqHNT
zI8MWEeTT+jpFTSR4Qt/Gdp5PbzTxfun38oOgT8WB3xJ1XvrRsxROPluBa1y7cVm
UcriTPzUtAhxb7MVGaTVwQ1zX1Wd+t0mYQVW8zRMK6ECAwEAAaNFMEMwEgYDVR0T
AQH/BAgwBgEB/wIBATAOBgNVHQ8BAf8EBAMCAUYwHQYDVR0OBBYEFLzBQbsbg8is
pyWorw6eP2ftJETsMA0GCSqGSIb3DQEBBQUAA4IBAQDXb48+km740mC/EE68jHts
QV9tAFJ2c0WhMUfn0quL1C7FCUu9Y2lq75Rw7knbi+Q+F+PL165pk9WKQ/Q8iW3/
E7DBy67uV6r/3PT+Ay4GemfOMWj+MKaJQD5+EBErnqNXglfYZvG6JQorHtz29OFb
GJ3/dICwhz/SFF2/Hxh8mpzGpRs5CPMpSD6sFc+MhK8JsWzpOCRIHGzStF47dyG0
fY7KVrPFmx46Fx6aoNOF4AS8rMNcVaYmlHGhEn546LK3e+UeapK8GN9haNrggbTs
Eg+Uruj2i6nbXOuVJkJAIpbx/KuPb2vy+NCbLoPekfufWzFyy0Cs8CSU9CeLeaH0
-----END CERTIFICATE-----
`)
	// RSA2048Key is a 2048-bit RSA key.
	RSA2048Key = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA6PbRgXHcujfHGyEbAKKtaLGndbFr6TcWlrejV9nf6OXVXSin
izCt/0RySusM0dKQN5ZFWNK5gBn+3wpP8H8m4/ysvw9aRAPQFRAC4CvEiwYm6F/C
42EzqaFIQQwitvqDVctDfutblr0VMhtj5AmBiw+BaMVGBjlxP5leuBpyaDpxVYXv
nBKNbkPNlq8ZV8c8N5uXR1ep5Qde+Sxln+l/OxPKqWBWncGvk8o43pSoc1PMjwxY
R5NP6OkVNJHhC38Z2nk9vNPF+6ffyg6BPxYHfEnVe+tGzFE4+W4FrXLtxWZRyuJM
/NS0CHFvsxUZpNXBDXNfVZ363SZhBVbzNEwroQIDAQABAoIBAQDLnr/rxlvJH+uV
mNADNC0hbvYRdqv9QbsqrQPGS5bb99cP//LBRCExFuBW/y9LTiHjlCK0yip8/zu4
M0k/ycNyTm2m6YJaJIMBhecdjOPTJ+NmBB1RpKoFQATpZfQJvtiAapNqIckB7e7S
xwH+VRi3NSxFKPhVhGupzSHvBJ6u3Yrx49kAX7CDVlRFAu7NGkDmbkN5gknxHAFt
qwd6uLIrUwPQ3OJoqleU8ASYzI1CdGqSFojl67bYmanXbtQoxYFqtwgkWucttzdl
hfpCOw+kiB1LGQI7RNnW3yAfII7QLsO+nVNQgtxMe7qWxxNCMDSEnhRzNsicVKp/
n+vDTu4pAoGBAP17l2Gz8ZY8RiS7Kmjkyt5ZaEHKva81L6fgnN9rL105up8hm+CB
paqLOKh0DQHcMiBDkrUwbVvvPp2oq8iu6Uui/mihlyEnlkM35PpV8HIqcvDFh+Jo
6lopjM635qLW9uHyQ4d+mF2V6NqvGv01dE30HJEqDmaMR3dTZ9OvbtHTAoGBAOtH
EG0ezQXOAQWqiAq771pDRx+k0M8P4lu2f0mLFB53M92dxS3/hYjQJvpNwvJdPXLP
jzfv04MjN6vW1X+pol8xpCHYCNlPmjWt+xW73mZVTLM74SNjYQ44v2x4pF9g9nng
rX44aM+LqKXO5zu9dWM9JuRCe17sP1ElF6knRPA7AoGACLPXjKkq4CeNmPE8EYHZ
XSzgoXGedYdz7WWOvTTm2WKD/7adrWWGFIbXGSFy2N+AcQ8g2EujVYavNaZ2z1sB
83DTHzB9CcxcIk6m89lDegfvDkkZ0zIa6aGHjglOR8TtkPBKVTqJbJ0a83cTjCHr
rkl1OZ6iA+9I/NXGOMRLH7UCgYEA30P1m4diCYMexzC3nnAPR7mWUboGiKfLJzr8
eV6ofeyiZEimZ+sV3emhQ1/tgi7m8/9xKiTEs6oE12Wr/lSMiAdEePVYGFgIv63V
Gh/IgZWqjl9hW0KgRG2ngZjOatBJtQh3utJu65zdMlMwbSlxrvXF5VANYNuRjkBD
vrpMGicCgYAnMLwEVnoW7yNpsz4KrvXMQigQ3zNMDYFZPlwNMRSo1zl8k4OAPVL1
U76uzbRNRlCGtKPKRwQhcSxrc6gNuCd84l1t1goCBvkQk/c0q2J/8YQi743OJLT6
1HttNHgxkzTDmn72TepqDq/eMCSWzuoN+fFDnJZdK88hiCgCTHIGrQ==
-----END RSA PRIVATE KEY-----
`)

	// RSA1024Cert is an RSA CA cert with a 1024-bit key, SHA256 signature algorithm, that is currently valid and expires in 2117.
	// This should not be a valid cert because the key is only 1024 bits.
	RSA1024Cert = []byte(`
-----BEGIN CERTIFICATE-----
MIIChzCCAfCgAwIBAgIJAK9Xim2q4NaMMA0GCSqGSIb3DQEBCwUAMGExCzAJBgNV
BAYTAlVTMQswCQYDVQQIEwJDQTEWMBQGA1UEBxMNU2FuIEZyYW5jaXNjbzEPMA0G
A1UEChMGRG9ja2VyMRwwGgYDVQQDExNTd2FybWtpdCBDQSBUZXN0aW5nMCAXDTE3
MDEyNjIzMTQ1MFoYDzIxMTcwMTAyMjMxNDUwWjBhMQswCQYDVQQGEwJVUzELMAkG
A1UECBMCQ0ExFjAUBgNVBAcTDVNhbiBGcmFuY2lzY28xDzANBgNVBAoTBkRvY2tl
cjEcMBoGA1UEAxMTU3dhcm1raXQgQ0EgVGVzdGluZzCBnzANBgkqhkiG9w0BAQEF
AAOBjQAwgYkCgYEAwJecFi5Sa4aaY5lRvZZbiDA9ETESO7xrIgVWM3OVvBFAb8k2
9CRkxSpalEp4Iguwl6i3liMXudFXpek8sVcqzZDbFeQ6GfPL2zQU7hLevvhutE1V
moj8L5khsdyhDLwLBLl8XCYNCq4WlJvzuK4vKcO6bRc+2hlpogmOWFwjfBECAwEA
AaNFMEMwEgYDVR0TAQH/BAgwBgEB/wIBATAOBgNVHQ8BAf8EBAMCAUYwHQYDVR0O
BBYEFEjeSZQwqag+zm7sh85i0H6saGojMA0GCSqGSIb3DQEBCwUAA4GBADwPil+v
LfLlEZS1DrNy1nwl6mQuekqkfduq0U7fmaH6fpGYGs4Dbxjf/WqjV34EspMW6CGS
TCb+9eeYDfGqvZkSUwtpnN1m/1H19+2PD86aPRDQgeRE7BOhU0jsxJ3mYWwacMPH
fvP9c4cDXwEPJ/ocj95Ps35snJTpzFAaG7hp
-----END CERTIFICATE-----
`)
	// RSA1024Key is a 1024-bit RSA key
	RSA1024Key = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDAl5wWLlJrhppjmVG9lluIMD0RMRI7vGsiBVYzc5W8EUBvyTb0
JGTFKlqUSngiC7CXqLeWIxe50Vel6TyxVyrNkNsV5DoZ88vbNBTuEt6++G60TVWa
iPwvmSGx3KEMvAsEuXxcJg0KrhaUm/O4ri8pw7ptFz7aGWmiCY5YXCN8EQIDAQAB
AoGBALYWIWLvRMmYp5uHN7sxzzSBtxrr9Ds6N2gg95EJtQXsoamO6kAFsKihFKaj
idVWjA23XGu8ng/3FxEr5VAeA75WMnd82XxGCDostRwufBU2N6O96MMAiTCEia5q
lttn7OE4kgW4tSrTODKM6utvkqmLyJJeFlPHgoEb0WI6L95hAkEA7x9xMjd5WFES
t/cloA4msaIVSDbzN9ql31Z9IP/0z6CexNj3pjdtRD+Ydj9dPIzeskoDseS2d0l2
RXX3Z9YYJQJBAM4vb5UxVY4qaCY/tS44tAf6vwIo0lzKHBd41+ubpefWL6C4lhd1
jLhmwY6dio7mzFfKeI5Xtdu6DXr0zClzSn0CQGLpaaRxB/O9TXXleJ3VXLIbrpv5
hu/ytKxGlWniFn0QHrykVwRdZwhVGhbHrSSPzMqJDTA3wDZln9OpsVY1XDUCQQCr
hL54B8A6MYDOQLUBrF3nPWnj6/2C/wZ7aCWGc8aBo6WfN65z+W+EfsaJUvjOg6R9
a4r6LnC0RoOsQzQLT0MpAkA7q59Eo9DwPuLz6GrGAKBaxYXXPOyx58yO4DAq0e32
anuVw1kAAKz5HYioZkBJpnpN5dXCHNC54ooq76cIGFpT
-----END RSA PRIVATE KEY-----
`)

	// ECDSA224Cert is an ECDSA curve-P224 CA cert with a SHA256 signature algorithm
	// that is current valid and expires in 2117. This should not be a valid cert because we only accept curve-P256,
	// curve-P385, and curve-P521 (the only keys cfssl will generate).
	ECDSA224Cert = []byte(`
-----BEGIN CERTIFICATE-----
MIIB7jCCAZugAwIBAgIJALF0a2jHg8P9MAoGCCqGSM49BAMCMGExCzAJBgNVBAYT
AlVTMQswCQYDVQQIEwJDQTEWMBQGA1UEBxMNU2FuIEZyYW5jaXNjbzEPMA0GA1UE
ChMGRG9ja2VyMRwwGgYDVQQDExNTd2FybWtpdCBDQSBUZXN0aW5nMCAXDTE3MDEy
NzAwMjg1MloYDzIxMTcwMTAzMDAyODUyWjBhMQswCQYDVQQGEwJVUzELMAkGA1UE
CBMCQ0ExFjAUBgNVBAcTDVNhbiBGcmFuY2lzY28xDzANBgNVBAoTBkRvY2tlcjEc
MBoGA1UEAxMTU3dhcm1raXQgQ0EgVGVzdGluZzBOMBAGByqGSM49AgEGBSuBBAAh
AzoABFseGAWIbCHKia0TN6tjJbzu4GOi6lqxitimkygWnxaROVo1sJ/61A0lmy7z
Z5nb3HRWfrDJYZbao0UwQzASBgNVHRMBAf8ECDAGAQH/AgEBMA4GA1UdDwEB/wQE
AwIBRjAdBgNVHQ4EFgQU93VkqOtp8QHVRh7qh22G+QsnO2QwCgYIKoZIzj0EAwID
QQAwPgIdAMD758a1UD/YBA/fc00XL5g+a6v3bt9ZiSwSifMCHQDu1/WD9JmCdjbB
UJrkTcIE8xDejpxjPooK1cLT
-----END CERTIFICATE-----
`)
	// ECDSA224Key is an ECDSA curve-P224 key.
	ECDSA224Key = []byte(`
-----BEGIN EC PRIVATE KEY-----
MGgCAQEEHK+OanuZ3Gqx7/xipRzOneQUUlc11AMavfj2d1qgBwYFK4EEACGhPAM6
AARbHhgFiGwhyomtEzerYyW87uBjoupasYrYppMoFp8WkTlaNbCf+tQNJZsu82eZ
29x0Vn6wyWGW2g==
-----END EC PRIVATE KEY-----
`)

	// ECDSA256SHA256Cert is an ECDSA curve-P256 CA cert with a SHA256 signature algorithm
	// that is current valid and expires in 2117. This is a valid cert because it has an accepted key length
	// and an accepted signature algorithm.
	ECDSA256SHA256Cert = []byte(`
-----BEGIN CERTIFICATE-----
MIICADCCAaagAwIBAgIJAOnbqU2SK/veMAoGCCqGSM49BAMCMGExCzAJBgNVBAYT
AlVTMQswCQYDVQQIEwJDQTEWMBQGA1UEBxMNU2FuIEZyYW5jaXNjbzEPMA0GA1UE
ChMGRG9ja2VyMRwwGgYDVQQDExNTd2FybWtpdCBDQSBUZXN0aW5nMCAXDTE3MDEy
NzAwMjcyNVoYDzIxMTcwMTAzMDAyNzI1WjBhMQswCQYDVQQGEwJVUzELMAkGA1UE
CBMCQ0ExFjAUBgNVBAcTDVNhbiBGcmFuY2lzY28xDzANBgNVBAoTBkRvY2tlcjEc
MBoGA1UEAxMTU3dhcm1raXQgQ0EgVGVzdGluZzBZMBMGByqGSM49AgEGCCqGSM49
AwEHA0IABHmyfgFJLu94IyPYeYv/laDUe6cXcZWZL62dW3tm61YUDRQb57zJxvaI
eHsd7KW0YwQEbOeh2Qo0Uab4+pgTsiWjRTBDMBIGA1UdEwEB/wQIMAYBAf8CAQEw
DgYDVR0PAQH/BAQDAgFGMB0GA1UdDgQWBBTcjpX4ZO+MWsSyKARyyRproJzAWjAK
BggqhkjOPQQDAgNIADBFAiAdIZG7qzr+vCSt6FnotFKOhRBpLw9vkq8O2kBNbPCy
4wIhANXcKDlG507bv5bOWYo92XDWuHd1FzyZfSLren9uFVfB
-----END CERTIFICATE-----
`)
	// ECDSA256SHA1Cert is an ECDSA curve-P256 CA cert with a SHA1 signature algorithm
	// that is current valid and expires in 2117. This should not be a valid cert because a SHA1 signature algorithm.
	ECDSA256SHA1Cert = []byte(`
-----BEGIN CERTIFICATE-----
MIIB/jCCAaWgAwIBAgIJAKGcB/unE+cZMAkGByqGSM49BAEwYTELMAkGA1UEBhMC
VVMxCzAJBgNVBAgTAkNBMRYwFAYDVQQHEw1TYW4gRnJhbmNpc2NvMQ8wDQYDVQQK
EwZEb2NrZXIxHDAaBgNVBAMTE1N3YXJta2l0IENBIFRlc3RpbmcwIBcNMTcwMTI3
MDAyNzQ0WhgPMjExNzAxMDMwMDI3NDRaMGExCzAJBgNVBAYTAlVTMQswCQYDVQQI
EwJDQTEWMBQGA1UEBxMNU2FuIEZyYW5jaXNjbzEPMA0GA1UEChMGRG9ja2VyMRww
GgYDVQQDExNTd2FybWtpdCBDQSBUZXN0aW5nMFkwEwYHKoZIzj0CAQYIKoZIzj0D
AQcDQgAEebJ+AUku73gjI9h5i/+VoNR7pxdxlZkvrZ1be2brVhQNFBvnvMnG9oh4
ex3spbRjBARs56HZCjRRpvj6mBOyJaNFMEMwEgYDVR0TAQH/BAgwBgEB/wIBATAO
BgNVHQ8BAf8EBAMCAUYwHQYDVR0OBBYEFNyOlfhk74xaxLIoBHLJGmugnMBaMAkG
ByqGSM49BAEDSAAwRQIgX90Mxm8eGW43u6ztz3ePHz9X8UEozx4311fyYwtsLTEC
IQC7EWwxn+xAzcHUzQ1INPrsmnuvladTumv5huhkARtlgg==
-----END CERTIFICATE-----
`)
	// ECDSA256Key is an ECDSA curve-P256 key.
	ECDSA256Key = []byte(`
-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIKXkvFfUcVbH9Uqxkdo4Obwc3RSJfEH2254sfqkx50xBoAoGCCqGSM49
AwEHoUQDQgAEebJ+AUku73gjI9h5i/+VoNR7pxdxlZkvrZ1be2brVhQNFBvnvMnG
9oh4ex3spbRjBARs56HZCjRRpvj6mBOyJQ==
-----END EC PRIVATE KEY-----
`)

	// DSA2048Cert is a DSA CA cert with a 2048 key, SHA1 hash, that is currently valid and expires in 2117
	// This should not be a valid cert because we do not accept DSA keys.
	DSA2048Cert = []byte(`
-----BEGIN CERTIFICATE-----
MIIEyTCCBIigAwIBAgIJANu4Tu71eD7AMAkGByqGSM44BAMwYTELMAkGA1UEBhMC
VVMxCzAJBgNVBAgTAkNBMRYwFAYDVQQHEw1TYW4gRnJhbmNpc2NvMQ8wDQYDVQQK
EwZEb2NrZXIxHDAaBgNVBAMTE1N3YXJta2l0IENBIFRlc3RpbmcwIBcNMTcwMTI2
MTgzNDQ2WhgPMjExNzAxMDIxODM0NDZaMGExCzAJBgNVBAYTAlVTMQswCQYDVQQI
EwJDQTEWMBQGA1UEBxMNU2FuIEZyYW5jaXNjbzEPMA0GA1UEChMGRG9ja2VyMRww
GgYDVQQDExNTd2FybWtpdCBDQSBUZXN0aW5nMIIDOjCCAi0GByqGSM44BAEwggIg
AoIBAQD0EIQuOHBiDsmKGxTe5Ck87A2J1kkFkHZzcg3kde3BMyfeP5r1ReDkXdYR
06r8e2De4Ymsu/B4p5qetiP0XcMO6fERyrBrGSxIANeJNM9ccRsfcnxnvSFIu1qk
LixSEQxE8wN4v/c7fyFZrtSxXly2CWxb4qlPIs/xoQs+s8pRuW/uFk18QjszYq96
cliIIAf1qNEqadnYRvHSX0Xn2J+PSW2aRXXr79C1AUNq/U/CVkMJ1RHq0jTwsxlA
3P6ofjhxW/rXY7uTZeeZBLLeU/sRugvRfiubWIkjl1h0frOk7S1sND5wZ6zCZORA
bEpd9yRsvPYKlUMnHy7oUGT/IF1tAhUA6SQtzdKO+BoiRLmJ29etE+KnLwMCggEA
YYEJJRA869RzyrCUxEOmOFumnPVWIrS0+SY/fdK6uxLDVhO5v0EKsx4f8rBS9PPA
L6/elbV/GYtnR5iKktx16X8Jeo2YT5madLamREkI/9C4x0+UKF6ETx+ttEkntdAv
d6H3tTJw0y9WOV+TyQpNl8PloqEHP2slpeUjXapfhia/kfKeKfR2rSAlnMyWeiHD
ANnAJn+dfoITSxHgyaao73fCMryPfmEK4ffNEVHd5SA1SUUeAmEqbTwDi0BD31w5
PU1kDthsbNYFEx3S7PThZeLL74xxNbjoMK4zTTueXFjLlhDr7YfZYzCGauxT/Cij
qSJxfojjLv4PGFgeoIiNwgOCAQUAAoIBAHFK5SqxjgLqmWcJERnkFxDWE3fcO9ow
lSHJXugzP5Uyv3+IYJ67J22QthsajrnSduCJ+TPgnGPkJHk+3zzFYKArNKOKC5si
MkUD8DBLhY23ieX01J34Ej+t/uQYge1zFaGNm3c1k3WuCTCsbYqJtn60sh50oG3q
lfeRiVFgDto5EraYG9AgtfPSSkeSFVxIBfu6Hy/ri5M9gYwsmVpHZFElCNCbCcnh
zeNosUe5DlYnCdeviY8y3GeIP7QctnFGCCNODOGTuAGoOYb0xSw7rLM1cNns5Xzh
iq4iRFElvjPuiYGAUAsSYqCGx7gt2TiWW4AWbCkZi3S86ppxeevI2OijRTBDMBIG
A1UdEwEB/wQIMAYBAf8CAQEwDgYDVR0PAQH/BAQDAgFGMB0GA1UdDgQWBBQfWhZE
rEu8JT69BxWWXujrVlrenTAJBgcqhkjOOAQDAzAAMC0CFQDVY5dfGv4GiM8HXqUM
Ve+sDSZ9OAIUd4Cznid6BdEVGyQop2PFd/48Ieo=
-----END CERTIFICATE-----
`)
	// DSA2048Key is a 2048-bit DSA key
	DSA2048Key = []byte(`
-----BEGIN DSA PRIVATE KEY-----
MIIDPgIBAAKCAQEA9BCELjhwYg7JihsU3uQpPOwNidZJBZB2c3IN5HXtwTMn3j+a
9UXg5F3WEdOq/Htg3uGJrLvweKeanrYj9F3DDunxEcqwaxksSADXiTTPXHEbH3J8
Z70hSLtapC4sUhEMRPMDeL/3O38hWa7UsV5ctglsW+KpTyLP8aELPrPKUblv7hZN
fEI7M2KvenJYiCAH9ajRKmnZ2Ebx0l9F59ifj0ltmkV16+/QtQFDav1PwlZDCdUR
6tI08LMZQNz+qH44cVv612O7k2XnmQSy3lP7EboL0X4rm1iJI5dYdH6zpO0tbDQ+
cGeswmTkQGxKXfckbLz2CpVDJx8u6FBk/yBdbQIVAOkkLc3SjvgaIkS5idvXrRPi
py8DAoIBAGGBCSUQPOvUc8qwlMRDpjhbppz1ViK0tPkmP33SursSw1YTub9BCrMe
H/KwUvTzwC+v3pW1fxmLZ0eYipLcdel/CXqNmE+ZmnS2pkRJCP/QuMdPlChehE8f
rbRJJ7XQL3eh97UycNMvVjlfk8kKTZfD5aKhBz9rJaXlI12qX4Ymv5Hynin0dq0g
JZzMlnohwwDZwCZ/nX6CE0sR4MmmqO93wjK8j35hCuH3zRFR3eUgNUlFHgJhKm08
A4tAQ99cOT1NZA7YbGzWBRMd0uz04WXiy++McTW46DCuM007nlxYy5YQ6+2H2WMw
hmrsU/woo6kicX6I4y7+DxhYHqCIjcICggEAcUrlKrGOAuqZZwkRGeQXENYTd9w7
2jCVIcle6DM/lTK/f4hgnrsnbZC2GxqOudJ24In5M+CcY+QkeT7fPMVgoCs0o4oL
myIyRQPwMEuFjbeJ5fTUnfgSP63+5BiB7XMVoY2bdzWTda4JMKxtiom2frSyHnSg
beqV95GJUWAO2jkStpgb0CC189JKR5IVXEgF+7ofL+uLkz2BjCyZWkdkUSUI0JsJ
yeHN42ixR7kOVicJ16+JjzLcZ4g/tBy2cUYII04M4ZO4Aag5hvTFLDusszVw2ezl
fOGKriJEUSW+M+6JgYBQCxJioIbHuC3ZOJZbgBZsKRmLdLzqmnF568jY6AIVAJ8Z
5HzoPpFuQiZ6/H/N6RYpQmAO
-----END DSA PRIVATE KEY-----
`)
	// ECDSACertChain contains 3 SHA256 curve P-256 certificates:  leaf, intermediate, and root
	// They all expire in 2117.  The leaf cert's OU is swarm-manager.
	ECDSACertChain = [][]byte{
		[]byte(`
-----BEGIN CERTIFICATE-----
MIIB3TCCAYOgAwIBAgIUG2izItTi/0YNpfdwUwo7UcjddawwCgYIKoZIzj0EAwIw
EjEQMA4GA1UEAxMHcm9vdENOMjAgFw0xNzAzMDEyMzA1MDBaGA8yMTE3MDIwNjAw
MDUwMFowKDEMMAoGA1UEChMDb3JnMQswCQYDVQQLEwJvdTELMAkGA1UEAxMCY24w
WTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATCVPwZBGYQ0SpeXahXzU8BB+ZBjdw9
WsKBa03qSic4O0qtUrLTQSvg2bWoKlo2fVe5g6Sl29gMm0912fTG5nHro4GeMIGb
MA4GA1UdDwEB/wQEAwIFoDAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIw
DAYDVR0TAQH/BAIwADAdBgNVHQ4EFgQU/hk9CSt3C+8+hVVe1+xTHdAYka4wHwYD
VR0jBBgwFoAU0qlzziAdvItofIcj5PK+SLIRngAwHAYDVR0RBBUwE4ICY26CDXN3
YXJtLW1hbmFnZXIwCgYIKoZIzj0EAwIDSAAwRQIhAIV+zZKA58KkkeV9lC7EgVjT
nXZuicOq8369KseHDSINAiAy8QKshS5XUHXFJi778Mclr2jvx88XnV2yYb7osJv4
Ew==
-----END CERTIFICATE-----
`),
		[]byte(`
-----BEGIN CERTIFICATE-----
MIIBizCCATCgAwIBAgIUcGcL0qGDloPcLE69t6X81DKiaZAwCgYIKoZIzj0EAwIw
ETEPMA0GA1UEAxMGcm9vdENOMCAXDTE3MDMwMjAwMDAwMFoYDzIxMTcwMjA2MDAw
MDAwWjASMRAwDgYDVQQDEwdyb290Q04yMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcD
QgAEL4g4/wWhZM/YfCk/zEXbmTIgaiNUsXrqexXGrsFeoxfojAEuA8tygI8mu45V
fNk16nzO4AfXMFBiChB9fPE1dKNjMGEwDgYDVR0PAQH/BAQDAgEGMA8GA1UdEwEB
/wQFMAMBAf8wHQYDVR0OBBYEFNKpc84gHbyLaHyHI+TyvkiyEZ4AMB8GA1UdIwQY
MBaAFGD5gOqAIojsuSKECZwWE5aeGDD9MAoGCCqGSM49BAMCA0kAMEYCIQDN10Lz
9mqWPOgqlpSboPf+VzC0HA1ZZI5wqETUKCK1wQIhANkepyJrCapiQ6Vuvc+qycuS
ZS16fmlAEKrBm2KgpZt2
-----END CERTIFICATE-----
`),
		[]byte(`
-----BEGIN CERTIFICATE-----
MIIBaDCCAQ6gAwIBAgIUfmVlMNH1dFyOjZHL18pw0ji9aTkwCgYIKoZIzj0EAwIw
ETEPMA0GA1UEAxMGcm9vdENOMCAXDTE3MDMwMjAwMDAwMFoYDzIxMTcwMjA2MDAw
MDAwWjARMQ8wDQYDVQQDEwZyb290Q04wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNC
AAT6NjQeSstS/gi2wN+AoWnMZaLfiBjpNSqryqEiPH03viwbtWMG9aCu7cU/3alJ
iIlmQl6Y3n3cFhiQV2dum+UUo0IwQDAOBgNVHQ8BAf8EBAMCAQYwDwYDVR0TAQH/
BAUwAwEB/zAdBgNVHQ4EFgQUYPmA6oAiiOy5IoQJnBYTlp4YMP0wCgYIKoZIzj0E
AwIDSAAwRQIgP8iV0PKFeQZey6j89ieI+IPucjfl8Hp1OLJbamrVEr8CIQD0PsI8
pMJFqD7k4votyNu3W82NrBSe+xyMgFqI5tfx4g==
-----END CERTIFICATE-----
`),
	}

	// ECDSACertChainKeys contains 3 SHA256 curve P-256 keys: corresponding, respectively,
	// to the certificates in ECDSACertChain
	ECDSACertChainKeys = [][]byte{
		[]byte(`
-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIN+BaGyxGLSgEDLjmQBHdL7JuuAIYlSGCwYS2CCUxMEOoAoGCCqGSM49
AwEHoUQDQgAEwlT8GQRmENEqXl2oV81PAQfmQY3cPVrCgWtN6konODtKrVKy00Er
4Nm1qCpaNn1XuYOkpdvYDJtPddn0xuZx6w==
-----END EC PRIVATE KEY-----
`),
		[]byte(`
-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIP7yNfaUImD76q1pfgx+8PYSq50zK1imh41SKFPzR5fioAoGCCqGSM49
AwEHoUQDQgAEL4g4/wWhZM/YfCk/zEXbmTIgaiNUsXrqexXGrsFeoxfojAEuA8ty
gI8mu45VfNk16nzO4AfXMFBiChB9fPE1dA==
-----END EC PRIVATE KEY-----
`),
		[]byte(`
-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIDIgEpCpn7wEEYt/hLT+NewO0lgBPBRk3A5nU4ASOShDoAoGCCqGSM49
AwEHoUQDQgAE+jY0HkrLUv4ItsDfgKFpzGWi34gY6TUqq8qhIjx9N74sG7VjBvWg
ru3FP92pSYiJZkJemN593BYYkFdnbpvlFA==
-----END EC PRIVATE KEY-----
`),
	}
)
