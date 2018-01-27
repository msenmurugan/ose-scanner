package fixchain

// Chain 1:
const googleLeaf = `-----BEGIN CERTIFICATE-----
MIIDITCCAoqgAwIBAgIQL9+89q6RUm0PmqPfQDQ+mjANBgkqhkiG9w0BAQUFADBM
MQswCQYDVQQGEwJaQTElMCMGA1UEChMcVGhhd3RlIENvbnN1bHRpbmcgKFB0eSkg
THRkLjEWMBQGA1UEAxMNVGhhd3RlIFNHQyBDQTAeFw0wOTEyMTgwMDAwMDBaFw0x
MTEyMTgyMzU5NTlaMGgxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpDYWxpZm9ybmlh
MRYwFAYDVQQHFA1Nb3VudGFpbiBWaWV3MRMwEQYDVQQKFApHb29nbGUgSW5jMRcw
FQYDVQQDFA53d3cuZ29vZ2xlLmNvbTCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkC
gYEA6PmGD5D6htffvXImttdEAoN4c9kCKO+IRTn7EOh8rqk41XXGOOsKFQebg+jN
gtXj9xVoRaELGYW84u+E593y17iYwqG7tcFR39SDAqc9BkJb4SLD3muFXxzW2k6L
05vuuWciKh0R73mkszeK9P4Y/bz5RiNQl/Os/CRGK1w7t0UCAwEAAaOB5zCB5DAM
BgNVHRMBAf8EAjAAMDYGA1UdHwQvMC0wK6ApoCeGJWh0dHA6Ly9jcmwudGhhd3Rl
LmNvbS9UaGF3dGVTR0NDQS5jcmwwKAYDVR0lBCEwHwYIKwYBBQUHAwEGCCsGAQUF
BwMCBglghkgBhvhCBAEwcgYIKwYBBQUHAQEEZjBkMCIGCCsGAQUFBzABhhZodHRw
Oi8vb2NzcC50aGF3dGUuY29tMD4GCCsGAQUFBzAChjJodHRwOi8vd3d3LnRoYXd0
ZS5jb20vcmVwb3NpdG9yeS9UaGF3dGVfU0dDX0NBLmNydDANBgkqhkiG9w0BAQUF
AAOBgQCfQ89bxFApsb/isJr/aiEdLRLDLE5a+RLizrmCUi3nHX4adpaQedEkUjh5
u2ONgJd8IyAPkU0Wueru9G2Jysa9zCRo1kNbzipYvzwY4OA8Ys+WAi0oR1A04Se6
z5nRUP8pJcA2NhUzUnC+MY+f6H/nEQyNv4SgQhqAibAxWEEHXw==
-----END CERTIFICATE-----
`

const thawteIntermediate = `-----BEGIN CERTIFICATE-----
MIIDIzCCAoygAwIBAgIEMAAAAjANBgkqhkiG9w0BAQUFADBfMQswCQYDVQQGEwJV
UzEXMBUGA1UEChMOVmVyaVNpZ24sIEluYy4xNzA1BgNVBAsTLkNsYXNzIDMgUHVi
bGljIFByaW1hcnkgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwHhcNMDQwNTEzMDAw
MDAwWhcNMTQwNTEyMjM1OTU5WjBMMQswCQYDVQQGEwJaQTElMCMGA1UEChMcVGhh
d3RlIENvbnN1bHRpbmcgKFB0eSkgTHRkLjEWMBQGA1UEAxMNVGhhd3RlIFNHQyBD
QTCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEA1NNn0I0Vf67NMf59HZGhPwtx
PKzMyGT7Y/wySweUvW+Aui/hBJPAM/wJMyPpC3QrccQDxtLN4i/1CWPN/0ilAL/g
5/OIty0y3pg25gqtAHvEZEo7hHUD8nCSfQ5i9SGraTaEMXWQ+L/HbIgbBpV8yeWo
3nWhLHpo39XKHIdYYBkCAwEAAaOB/jCB+zASBgNVHRMBAf8ECDAGAQH/AgEAMAsG
A1UdDwQEAwIBBjARBglghkgBhvhCAQEEBAMCAQYwKAYDVR0RBCEwH6QdMBsxGTAX
BgNVBAMTEFByaXZhdGVMYWJlbDMtMTUwMQYDVR0fBCowKDAmoCSgIoYgaHR0cDov
L2NybC52ZXJpc2lnbi5jb20vcGNhMy5jcmwwMgYIKwYBBQUHAQEEJjAkMCIGCCsG
AQUFBzABhhZodHRwOi8vb2NzcC50aGF3dGUuY29tMDQGA1UdJQQtMCsGCCsGAQUF
BwMBBggrBgEFBQcDAgYJYIZIAYb4QgQBBgpghkgBhvhFAQgBMA0GCSqGSIb3DQEB
BQUAA4GBAFWsY+reod3SkF+fC852vhNRj5PZBSvIG3dLrWlQoe7e3P3bB+noOZTc
q3J5Lwa/q4FwxKjt6lM07e8eU9kGx1Yr0Vz00YqOtCuxN5BICEIlxT6Ky3/rbwTR
bcV0oveifHtgPHfNDs5IAn8BL7abN+AqKjbc1YXWrOU/VG+WHgWv
-----END CERTIFICATE-----
`

const verisignRoot = `-----BEGIN CERTIFICATE-----
MIICPDCCAaUCEHC65B0Q2Sk0tjjKewPMur8wDQYJKoZIhvcNAQECBQAwXzELMAkG
A1UEBhMCVVMxFzAVBgNVBAoTDlZlcmlTaWduLCBJbmMuMTcwNQYDVQQLEy5DbGFz
cyAzIFB1YmxpYyBQcmltYXJ5IENlcnRpZmljYXRpb24gQXV0aG9yaXR5MB4XDTk2
MDEyOTAwMDAwMFoXDTI4MDgwMTIzNTk1OVowXzELMAkGA1UEBhMCVVMxFzAVBgNV
BAoTDlZlcmlTaWduLCBJbmMuMTcwNQYDVQQLEy5DbGFzcyAzIFB1YmxpYyBQcmlt
YXJ5IENlcnRpZmljYXRpb24gQXV0aG9yaXR5MIGfMA0GCSqGSIb3DQEBAQUAA4GN
ADCBiQKBgQDJXFme8huKARS0EN8EQNvjV69qRUCPhAwL0TPZ2RHP7gJYHyX3KqhE
BarsAx94f56TuZoAqiN91qyFomNFx3InzPRMxnVx0jnvT0Lwdd8KkMaOIG+YD/is
I19wKTakyYbnsZogy1Olhec9vn2a/iRFM9x2Fe0PonFkTGUugWhFpwIDAQABMA0G
CSqGSIb3DQEBAgUAA4GBALtMEivPLCYATxQT3ab7/AoRhIzzKBxnki98tsX63/Do
lbwdj2wsqFHMc9ikwFPwTtYmwHYBV4GSXiHx0bH/59AhWM1pF+NEHJwZRDmJXNyc
AA9WjQKZ7aKQRUzkuxCkPfAyAw7xzvjoyVGM5mKf5p/AfbdynMk2OmufTqj/ZA1k
-----END CERTIFICATE-----
`

// Chain 2:
const megaLeaf = `-----BEGIN CERTIFICATE-----
MIIFOjCCBCKgAwIBAgIQWYE8Dup170kZ+k11Lg51OjANBgkqhkiG9w0BAQUFADBy
MQswCQYDVQQGEwJHQjEbMBkGA1UECBMSR3JlYXRlciBNYW5jaGVzdGVyMRAwDgYD
VQQHEwdTYWxmb3JkMRowGAYDVQQKExFDT01PRE8gQ0EgTGltaXRlZDEYMBYGA1UE
AxMPRXNzZW50aWFsU1NMIENBMB4XDTEyMTIxNDAwMDAwMFoXDTE0MTIxNDIzNTk1
OVowfzEhMB8GA1UECxMYRG9tYWluIENvbnRyb2wgVmFsaWRhdGVkMS4wLAYDVQQL
EyVIb3N0ZWQgYnkgSW5zdHJhIENvcnBvcmF0aW9uIFB0eS4gTFREMRUwEwYDVQQL
EwxFc3NlbnRpYWxTU0wxEzARBgNVBAMTCm1lZ2EuY28ubnowggEiMA0GCSqGSIb3
DQEBAQUAA4IBDwAwggEKAoIBAQDcxMCClae8BQIaJHBUIVttlLvhbK4XhXPk3RQ3
G5XA6tLZMBQ33l3F9knYJ0YErXtr8IdfYoulRQFmKFMJl9GtWyg4cGQi2Rcr5VN5
S5dA1vu4oyJBxE9fPELcK6Yz1vqaf+n6za+mYTiQYKggVdS8/s8hmNuXP9Zk1pIn
+q0pGsf8NAcSHMJgLqPQrTDw+zae4V03DvcYfNKjuno88d2226ld7MAmQZ7uRNsI
/CnkdelVs+akZsXf0szefSqMJlf08SY32t2jj4Ra7RApVYxOftD9nij/aLfuqOU6
ow6IgIcIG2ZvXLZwK87c5fxL7UAsTTV+M1sVv8jA33V2oKLhAgMBAAGjggG9MIIB
uTAfBgNVHSMEGDAWgBTay+qtWwhdzP/8JlTOSeVVxjj0+DAdBgNVHQ4EFgQUmP9l
6zhyrZ06Qj4zogt+6LKFk4AwDgYDVR0PAQH/BAQDAgWgMAwGA1UdEwEB/wQCMAAw
NAYDVR0lBC0wKwYIKwYBBQUHAwEGCCsGAQUFBwMCBgorBgEEAYI3CgMDBglghkgB
hvhCBAEwTwYDVR0gBEgwRjA6BgsrBgEEAbIxAQICBzArMCkGCCsGAQUFBwIBFh1o
dHRwczovL3NlY3VyZS5jb21vZG8uY29tL0NQUzAIBgZngQwBAgEwOwYDVR0fBDQw
MjAwoC6gLIYqaHR0cDovL2NybC5jb21vZG9jYS5jb20vRXNzZW50aWFsU1NMQ0Eu
Y3JsMG4GCCsGAQUFBwEBBGIwYDA4BggrBgEFBQcwAoYsaHR0cDovL2NydC5jb21v
ZG9jYS5jb20vRXNzZW50aWFsU1NMQ0FfMi5jcnQwJAYIKwYBBQUHMAGGGGh0dHA6
Ly9vY3NwLmNvbW9kb2NhLmNvbTAlBgNVHREEHjAcggptZWdhLmNvLm56gg53d3cu
bWVnYS5jby5uejANBgkqhkiG9w0BAQUFAAOCAQEAcYhrsPSvDuwihMOh0ZmRpbOE
Gw6LqKgLNTmaYUPQhzi2cyIjhUhNvugXQQlP5f0lp5j8cixmArafg1dTn4kQGgD3
ivtuhBTgKO1VYB/VRoAt6Lmswg3YqyiS7JiLDZxjoV7KoS5xdiaINfHDUaBBY4ZH
j2BUlPniNBjCqXe/HndUTVUewlxbVps9FyCmH+C4o9DWzdGBzDpCkcmo5nM+cp7q
ZhTIFTvZfo3zGuBoyu8BzuopCJcFRm3cRiXkpI7iOMUIixO1szkJS6WpL1sKdT73
UXp08U0LBqoqG130FbzEJBBV3ixbvY6BWMHoCWuaoF12KJnC5kHt2RoWAAgMXA==
-----END CERTIFICATE-----
`

const comodoIntermediate = `-----BEGIN CERTIFICATE-----
MIIFAzCCA+ugAwIBAgIQGLLLuqME8aAPwfLzJkYqSjANBgkqhkiG9w0BAQUFADCB
gTELMAkGA1UEBhMCR0IxGzAZBgNVBAgTEkdyZWF0ZXIgTWFuY2hlc3RlcjEQMA4G
A1UEBxMHU2FsZm9yZDEaMBgGA1UEChMRQ09NT0RPIENBIExpbWl0ZWQxJzAlBgNV
BAMTHkNPTU9ETyBDZXJ0aWZpY2F0aW9uIEF1dGhvcml0eTAeFw0wNjEyMDEwMDAw
MDBaFw0xOTEyMzEyMzU5NTlaMHIxCzAJBgNVBAYTAkdCMRswGQYDVQQIExJHcmVh
dGVyIE1hbmNoZXN0ZXIxEDAOBgNVBAcTB1NhbGZvcmQxGjAYBgNVBAoTEUNPTU9E
TyBDQSBMaW1pdGVkMRgwFgYDVQQDEw9Fc3NlbnRpYWxTU0wgQ0EwggEiMA0GCSqG
SIb3DQEBAQUAA4IBDwAwggEKAoIBAQCt8AiwcsargxIxF3CJhakgEtSYau2A1NHf
5I5ZLdOWIY120j8YC0YZYwvHIPPlC92AGvFaoL0dds23Izp0XmEbdaqb1IX04XiR
0y3hr/yYLgbSeT1awB8hLRyuIVPGOqchfr7tZ291HRqfalsGs2rjsQuqag7nbWzD
ypWMN84hHzWQfdvaGlyoiBSyD8gSIF/F03/o4Tjg27z5H6Gq1huQByH6RSRQXScq
oChBRVt9vKCiL6qbfltTxfEFFld+Edc7tNkBdtzffRDPUanlOPJ7FAB1WfnwWdsX
Pvev5gItpHnBXaIcw5rIp6gLSApqLn8tl2X2xQScRMiZln5+pN0vAgMBAAGjggGD
MIIBfzAfBgNVHSMEGDAWgBQLWOWLxkwVN6RAqTCpIb5HNlpW/zAdBgNVHQ4EFgQU
2svqrVsIXcz//CZUzknlVcY49PgwDgYDVR0PAQH/BAQDAgEGMBIGA1UdEwEB/wQI
MAYBAf8CAQAwIAYDVR0lBBkwFwYKKwYBBAGCNwoDAwYJYIZIAYb4QgQBMD4GA1Ud
IAQ3MDUwMwYEVR0gADArMCkGCCsGAQUFBwIBFh1odHRwczovL3NlY3VyZS5jb21v
ZG8uY29tL0NQUzBJBgNVHR8EQjBAMD6gPKA6hjhodHRwOi8vY3JsLmNvbW9kb2Nh
LmNvbS9DT01PRE9DZXJ0aWZpY2F0aW9uQXV0aG9yaXR5LmNybDBsBggrBgEFBQcB
AQRgMF4wNgYIKwYBBQUHMAKGKmh0dHA6Ly9jcnQuY29tb2RvY2EuY29tL0NvbW9k
b1VUTlNHQ0NBLmNydDAkBggrBgEFBQcwAYYYaHR0cDovL29jc3AuY29tb2RvY2Eu
Y29tMA0GCSqGSIb3DQEBBQUAA4IBAQAtlzR6QDLqcJcvgTtLeRJ3rvuq1xqo2l/z
odueTZbLN3qo6u6bldudu+Ennv1F7Q5Slqz0J790qpL0pcRDAB8OtXj5isWMcL2a
ejGjKdBZa0wztSz4iw+SY1dWrCRnilsvKcKxudokxeRiDn55w/65g+onO7wdQ7Vu
F6r7yJiIatnyfKH2cboZT7g440LX8NqxwCPf3dfxp+0Jj1agq8MLy6SSgIGSH6lv
+Wwz3D5XxqfyH8wqfOQsTEZf6/Nh9yvENZ+NWPU6g0QO2JOsTGvMd/QDzczc4BxL
XSXaPV7Od4rhPsbXlM1wSTz/Dr0ISKvlUhQVnQ6cGodWaK2cCQBk
-----END CERTIFICATE-----
`

const comodoRoot = `-----BEGIN CERTIFICATE-----
MIIEHTCCAwWgAwIBAgIQToEtioJl4AsC7j41AkblPTANBgkqhkiG9w0BAQUFADCB
gTELMAkGA1UEBhMCR0IxGzAZBgNVBAgTEkdyZWF0ZXIgTWFuY2hlc3RlcjEQMA4G
A1UEBxMHU2FsZm9yZDEaMBgGA1UEChMRQ09NT0RPIENBIExpbWl0ZWQxJzAlBgNV
BAMTHkNPTU9ETyBDZXJ0aWZpY2F0aW9uIEF1dGhvcml0eTAeFw0wNjEyMDEwMDAw
MDBaFw0yOTEyMzEyMzU5NTlaMIGBMQswCQYDVQQGEwJHQjEbMBkGA1UECBMSR3Jl
YXRlciBNYW5jaGVzdGVyMRAwDgYDVQQHEwdTYWxmb3JkMRowGAYDVQQKExFDT01P
RE8gQ0EgTGltaXRlZDEnMCUGA1UEAxMeQ09NT0RPIENlcnRpZmljYXRpb24gQXV0
aG9yaXR5MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0ECLi3LjkRv3
UcEbVASY06m/weaKXTuH+7uIzg3jLz8GlvCiKVCZrts7oVewdFFxze1CkU1B/qnI
2GqGd0S7WWaXUF601CxwRM/aN5VCaTwwxHGzUvAhTaHYujl8HJ6jJJ3ygxaYqhZ8
Q5sVW7euNJH+1GImGEaaP+vB+fGQV+useg2L23IwambV4EajcNxo2f8ESIl33rXp
+2dtQem8Ob0y2WIC8bGoPW43nOIv4tOiJovGuFVDiOEjPqXSJDlqR6sA1KGzqSX+
DT+nHbrTUcELpNqsOO9VUCQFZUaTNE8tja3G1CEZ0o7KBWFxB3NH5YoZEr0ETc5O
nKVIrLsm9wIDAQABo4GOMIGLMB0GA1UdDgQWBBQLWOWLxkwVN6RAqTCpIb5HNlpW
/zAOBgNVHQ8BAf8EBAMCAQYwDwYDVR0TAQH/BAUwAwEB/zBJBgNVHR8EQjBAMD6g
PKA6hjhodHRwOi8vY3JsLmNvbW9kb2NhLmNvbS9DT01PRE9DZXJ0aWZpY2F0aW9u
QXV0aG9yaXR5LmNybDANBgkqhkiG9w0BAQUFAAOCAQEAPpiem/Yb6dc5t3iuHXIY
SdOH5EOC6z/JqvWote9VfCFSZfnVDeFs9D6Mk3ORLgLETgdxb8CPOGEIqB6BCsAv
IC9Bi5HcSEW88cbeunZrM8gALTFGTO3nnc+IlP8zwFboJIYmuNg4ON8qa90SzMc/
RxdMosIGlgnW2/4/PEZB31jiVg88O8EckzXZOFKs7sjsLjBOlDW0JB9LeGna8gI4
zJVSk/BwJVmcIGfE7vmLV2H0knZ9P4SNVbfo5azV8fUZVqZa+5Acr5Pr5RzUZ5dd
BA6+C4OmF4O5MBKgxTMVBbkN+8cFduPYSo38NBejxiEovjBFMR7HeL5YYTisO+IB
ZQ==
-----END CERTIFICATE-----
`

// Chain 3:

const testLeaf = `-----BEGIN CERTIFICATE-----
MIIDrzCCApegAwIBAgIJAJlJNev65WpnMA0GCSqGSIb3DQEBCwUAMFoxCzAJBgNV
BAYTAlVLMRAwDgYDVQQIEwdFbmdsYW5kMQ8wDQYDVQQHEwZMb25kb24xEDAOBgNV
BAoTB0V4YW1wbGUxFjAUBgNVBAMTDUludGVybWVkaWF0ZTIwHhcNMTYwMjI5MTY1
MjAxWhcNMTYwMzMwMTY1MjAxWjBRMQswCQYDVQQGEwJVSzEQMA4GA1UECBMHRW5n
bGFuZDEPMA0GA1UEBxMGTG9uZG9uMRAwDgYDVQQKEwdFeGFtcGxlMQ0wCwYDVQQD
EwRMZWFmMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEApNsqW+/vRclm
Kbu/TcE6wfORzL8Nv9mGIsos94x1WrMOoRKMp3Slvczg9fGZsQNKrixW/38ZyUqW
KzIh/PKH4L3hgSQFsomZ7C1zZ4yiH+zZTf+fjPRX86q8hf1cOhP0YdMvYMmSfjHG
I7UDC9AAz6f8AmvK6CdNQaMlikZpfolWxQUsqhajVrHNv4LhjXFS/Yyi23D2qDpG
0e3kgQ+54OlWb3Gg2hsBZ6ddILukzLgkctYzHLtUv4EdAzol79yHdx+H7GxxF7c5
Ldju1H+b+Xc/dEBSKRCjEuJhXXmNlutfQPE6JzA9sr0wAXJBt/6n2UCEEZld4+XM
CtYaCxMrSwIDAQABo4GAMH4wDAYDVR0TBAUwAwEB/zAdBgNVHQ4EFgQUppd9scFX
d2jUJEeUHA87ttmyOjQwRAYIKwYBBQUHAQEEODA2MDQGCCsGAQUFBzAChihodHRw
Oi8vd3d3LmV4YW1wbGUuY29tL2ludGVybWVkaWF0ZTIuY3J0MAkGA1UdEQQCMAAw
DQYJKoZIhvcNAQELBQADggEBAIrtR+O1H20nUqxnBIS8/efGH7YUBpKGGbs3CmSW
7+IUke1VcDO3gNMjW7A/UxDM+1GM0MYD0/Pmlnk3/Q4TLZDpkAbk6lU5A/PVLqyE
5maPmwA+uIL3So9ivoCbIqbK/38g0Gqvdvq3yafH/60iodBAokr7r5iY/HmMBp4y
8PpQsZpx16XJWm0mkzvUx8OBS0MD/mnoCWE5i3Q9FT+KEEByZEyxj+2aqx3bYPhl
bqdbmwN5ZqYKt4lvJHJbjZb4gKKEEaFcBNJjZoSUDWnE+I1ZwmaXzkwwqfFEam1v
FSv/+x5C+55ylyX9E+S3UoH7wVswX/iT08hKL4IInmPoYQE=
-----END CERTIFICATE-----
`

const testIntermediate2 = `-----BEGIN CERTIFICATE-----
MIIDuDCCAqCgAwIBAgIJALmEjiSeHoB7MA0GCSqGSIb3DQEBCwUAMFoxCzAJBgNV
BAYTAlVLMRAwDgYDVQQIDAdFbmdsYW5kMQ8wDQYDVQQHDAZMb25kb24xEDAOBgNV
BAoMB0V4YW1wbGUxFjAUBgNVBAMMDUludGVybWVkaWF0ZTEwHhcNMTYwMjI5MTY0
NzQ1WhcNMTYwMzMwMTY0NzQ1WjBaMQswCQYDVQQGEwJVSzEQMA4GA1UECBMHRW5n
bGFuZDEPMA0GA1UEBxMGTG9uZG9uMRAwDgYDVQQKEwdFeGFtcGxlMRYwFAYDVQQD
Ew1JbnRlcm1lZGlhdGUyMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA
ukLNzAPup6wt5e4pGsZ1na6VIj7sx74WaenW2RNp2NrxTQIGPugU2HET8PPau/ch
VQNCtKLq8Wx9Kslq/aqn2e00jhxxrQmVc1fPwSnpOudhwbU6hvUkjYVtZ36BnGdp
oWw6HNVeemaBYFOtH6RXo/KtD7A+YLbfc5am//iMCdZI2oBeTLEToB3q21p/0PKm
pIKGZrPJnFdzSPlVkcDr/Lz8a1UCPYQW4zjPlYZjp9wDWpR7E7Fwla1RSFwBTnu1
xTVJnK7DU/i2A3JV9YLx63f9rBwwQdFBRWxr67GVeU6L7j43k/8H8CKCoOkk9Gf/
NP2pI6ZRh+toX4EzFaRL4QIDAQABo4GAMH4wDAYDVR0TBAUwAwEB/zAdBgNVHQ4E
FgQUSxliX4GqBki8rtWTgWYVp/1FXc0wRAYIKwYBBQUHAQEEODA2MDQGCCsGAQUF
BzAChihodHRwOi8vd3d3LmV4YW1wbGUuY29tL2ludGVybWVkaWF0ZTEuY3J0MAkG
A1UdEQQCMAAwDQYJKoZIhvcNAQELBQADggEBAINOzqDqNRgbWBEliCMQffLlJCAR
ypQ6U/jqeUlfZ2VKgPYo2wyloaMgOFKtaHUeOjiiQ1YJEgtP1BlyRmlazRi7iqWI
hvTtWom/8hyWG6AyN0tA4yK0+R4+OxMKDrNXU/C9W2p/yoI/fftU4m7QirqAW0ow
jzgjd3+M+eiIDDxBIwLLPPJGKXTqFBQ6U0LPTfNJrX1IzeSUTpHO7uoD/OWAlcC5
6LZxbPrgl3qb2bExavvXNDm4WkZwfG3iodi27FRIW8dydePbE/Ism1AADByMdzZ9
boo602kOaxt0kkvfgyUkXdRMlAhlzgTyGJmY69tZmvcWlsh7Rm+R5+fQeMI=
-----END CERTIFICATE-----
`

const testIntermediate1 = `-----BEGIN CERTIFICATE-----
MIIDkDCCAnigAwIBAgIJAK65INis8Pe8MA0GCSqGSIb3DQEBCwUAMD4xCzAJBgNV
BAYTAlVLMRAwDgYDVQQIEwdFbmdsYW5kMRAwDgYDVQQKEwdFeGFtcGxlMQswCQYD
VQQDEwJDQTAeFw0xNjAyMjkxNjM2MDdaFw0xNjAzMzAxNjM2MDdaMFoxCzAJBgNV
BAYTAlVLMRAwDgYDVQQIDAdFbmdsYW5kMQ8wDQYDVQQHDAZMb25kb24xEDAOBgNV
BAoMB0V4YW1wbGUxFjAUBgNVBAMMDUludGVybWVkaWF0ZTEwggEiMA0GCSqGSIb3
DQEBAQUAA4IBDwAwggEKAoIBAQDNibsPgy2npVpATi/JrOEsbjmBBRgLDXp8tA5T
JTl5YGn4K/9eltl12JvnXh5EgBxl8uTQwgaBX/+IY6BdXOcldfKPIBIO04gU9JtG
DnvW0638T2ujmwyDkpKZ4yTqugHo5nteLLDzt1iOMNeAZud0y5Lwxql4JIFI9eat
8C3JUESx9A8ZBNvAbcd38KBIoJAf67V5r2R3maqA/krxczVmUpf+xDuzkcEWDlgX
HfTC0mH2y2jsuPx4eepM7h+oPf0n/Nsoln4H+KiLL9l2bt5AtT1qdapQbcU1+GrU
I8Kqgvbe0XNlmZ6XYkQrprGcZTGmIsHNMWW+E6oyJni2O5TpAgMBAAGjdTBzMAwG
A1UdEwQFMAMBAf8wHQYDVR0OBBYEFL/b9VYR33yOBrw5QONVjHhb8+O5MDkGCCsG
AQUFBwEBBC0wKzApBggrBgEFBQcwAoYdaHR0cDovL3d3dy5leGFtcGxlLmNvbS9j
YS5jcnQwCQYDVR0RBAIwADANBgkqhkiG9w0BAQsFAAOCAQEAvwEKifIQX/Ua9PwF
J8i6tf4IfMaP1ra46n9uwwC3YPsQF8MjLys+mcaM899lroaPIRwy/FrRfx4lLsHm
EwN3jNgvRPDgUpGF7qsC2HcKUpvAzQoHQy/JN8Gy2StywtzvlaswR8qKR32zKoUc
y95NwaGzg4HxK3lNpj0Vorus6VED8GLgbiewAmFVlXd/pWrv8t1zzCyK780zetgD
R4xC6VVsV+6V/1MqiuVtYHUWJj6MBAWOXPnHeZB6BDjuQh7MdQr7Bx4EOzZnnzny
4hTZOFVQeV/7jx/NKQoh/oEyGBqNW8R5umynowG/SHFQGNKpyvzQvZr6enPbrHEI
pLLzxA==
-----END CERTIFICATE-----
`

const testRoot = `-----BEGIN CERTIFICATE-----
MIIDOTCCAiGgAwIBAgIJAJWqHxvK6nAQMA0GCSqGSIb3DQEBBQUAMD4xCzAJBgNV
BAYTAlVLMRAwDgYDVQQIEwdFbmdsYW5kMRAwDgYDVQQKEwdFeGFtcGxlMQswCQYD
VQQDEwJDQTAeFw0xNjAyMjkxNTU1MDBaFw0xNzAyMjgxNTU1MDBaMD4xCzAJBgNV
BAYTAlVLMRAwDgYDVQQIEwdFbmdsYW5kMRAwDgYDVQQKEwdFeGFtcGxlMQswCQYD
VQQDEwJDQTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBANKk/zKjG6Z4
Cux2o7v0HmkRuzMmn0jXizpH9+9g0MKCcIfbUK05YpUUX09Ij3u4t9zrbola+drt
0u5DV9FujtD3cjwOOxvSyGWzzqoI2uh5pOFDkTaEtBzvLEB9T1wsrvmLXy0IMEhu
QL8B33SY+7T5cuBPSIX0uipRJDz+1k+ESwfXhXz4uS1Y1K7GD+BxlPVrh71WxA5r
f+hHsDE9f2LvCTI156dJHtBVSzAFxq/Kl2r+fofFCKUyBo9wNgL3j4t/jxbkEwA+
YDpqZ1YFYSFD/3tn/vEF0KmdmW9L3zrCli1WWtAOkl7oVCgPiXs9mfjZqYj4+0la
hJhae79jS+sCAwEAAaM6MDgwHQYDVR0OBBYEFP+7oT2emG1zF5bhXfVyrUWCymEj
MAwGA1UdEwQFMAMBAf8wCQYDVR0RBAIwADANBgkqhkiG9w0BAQUFAAOCAQEAfbyn
lejnhbuiR5s7ENw1aMkod3FnZhhAIhCHncqsAJ1XPqaVvZxgrY3Rrxudp9rA9Gwn
ZJoPqLOLhcWzLkSpPQ1w43HOk19Ok9UGRsSpkHlPbTac3wzcjKEbBpPONgoin80/
ZGMvti8uvkZH8qqWsvmRrq4pDEK1h2eeF0ayF349evdKyFB2yfWV5dUrBXin28vf
AcQIk20/eb9bZ3KMPCa3dGjzDaQDAuutjS/XkRRkmMnp3q3ZQKlE8Fc+bzoP274Q
X2+I8S85sfOxJOR+CF6CY7IND2BlnUFCOXvIk2V8tX1lbP3Lv6Ukuz+AlQ7Whn5T
qiVjrghWwV/u6h3HSA==
-----END CERTIFICATE-----
`

// Chain 4: Contains a loop: C signed by B signed by A signed by B etc

const testC = `-----BEGIN CERTIFICATE-----
MIIDUjCCAjqgAwIBAgIJANYuoGKAZHTEMA0GCSqGSIb3DQEBBQUAMD0xCzAJBgNV
BAYTAlVLMRAwDgYDVQQIDAdFbmdsYW5kMRAwDgYDVQQKDAdFeGFtcGxlMQowCAYD
VQQDDAFCMB4XDTE2MDMwMTE2MjE1NFoXDTE3MDMwMTE2MjE1NFowPTELMAkGA1UE
BhMCVUsxEDAOBgNVBAgMB0VuZ2xhbmQxEDAOBgNVBAoMB0V4YW1wbGUxCjAIBgNV
BAMMAUMwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQC6Z4S+7d7dkUDH
fOY1AALL2MAg2PqjleuCk4fvsRz5zyVSAMH7R1MAI5atP5RrEwWdMxlXQO1QCaFf
j5P8uNzvCyVaryHkmObSHzgC1+NKpGXNOf1c8VV4XHZn4oLg8vtN1uMjO4Sv3tc6
MaAqja9EkgIqO6HoN5bhnyCpNJfyZDZEYtKamZqxLo0nABQ3kSlZd53ThqNaI26y
u3yKlWC/WHctYbiFMBx3GVzTIGjKHlHlYDs+nqF1KtV1Z46pQsIDXsWb/AJ/Je9c
JsE0JaPS86SRiQTf6/uyNNPLW1qteAe3yEl3Kc0hMunxDNbmRbKR8NO9xrHT5jFT
e9FhFhsBAgMBAAGjVTBTMAwGA1UdEwQFMAMBAf8wOAYIKwYBBQUHAQEELDAqMCgG
CCsGAQUFBzAChhxodHRwOi8vd3d3LmV4YW1wbGUuY29tL2IuY3J0MAkGA1UdEQQC
MAAwDQYJKoZIhvcNAQEFBQADggEBAI3DU3KzKfFEsHp1Ie+kPJeHZWOxk5Ti5L+G
M2OoIc7ht8qO1ty1PE0XgaZ+wjpiknWRj+RcDrETwB7iINatQhmvWkEhH6yl2GqJ
IbucsNouTCYxfP1K9kT4/XuXh0Jg7Q1cjHpQfcAazfr68XRLhcZv5cQmDBDwjUq8
OM1Xmp+BmYO+7XDM3g2j9d65xblgDFdWS4wM2pp63C3ywkuetXw/Yj2Pok/j+97D
kGA8psFaxN/4vtHVzJg6YX80bNvURFNMoTjulxqhLvP8UqZidI7kU+V/JBRBZGvH
ANWsKuJ5uadPBSat63pkIVTFsKwms1Vnfw39QDQi18HOp0zBzR0=
-----END CERTIFICATE-----
`

const testB = `-----BEGIN CERTIFICATE-----
MIIDUjCCAjqgAwIBAgIJAJWqHxvK6nASMA0GCSqGSIb3DQEBBQUAMD0xCzAJBgNV
BAYTAlVLMRAwDgYDVQQIDAdFbmdsYW5kMRAwDgYDVQQKDAdFeGFtcGxlMQowCAYD
VQQDDAFBMB4XDTE2MDMwMTE1MTMyMFoXDTE3MDMwMTE1MTMyMFowPTELMAkGA1UE
BhMCVUsxEDAOBgNVBAgMB0VuZ2xhbmQxEDAOBgNVBAoMB0V4YW1wbGUxCjAIBgNV
BAMMAUIwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDgQutw5xEEqBqn
1aS/Ye2xSJOV3RtOeZ8aEDOOrzaU/VL39Kk46+5tb8KbTDyri8HUfGe9JlKYPaIG
lxs6QBq0watILRgypzgKLkqmBprTQIzXpgWPkUwh2Hn6cR/8YMSbObxTXZmm6ELd
GG/dn2LfchzB94/JEVs4FZpVXRlfL7Wp8L/eCzmOMuN8eSyXEJWCSiWezjclh14V
znSZ3pmtxsdiBJcBVscVfsYyAJ/Zznpej5orbO7zn9kxxRNrVTlcW5twgnl07A6w
Tav2yWdNnFJZJYRYha4x9dsz5ZutCsiYQ9ZJs7MxYVRkGtn+cf92PqICbr25XY8O
RwPZU/FjAgMBAAGjVTBTMAwGA1UdEwQFMAMBAf8wOAYIKwYBBQUHAQEELDAqMCgG
CCsGAQUFBzAChhxodHRwOi8vd3d3LmV4YW1wbGUuY29tL2EuY3J0MAkGA1UdEQQC
MAAwDQYJKoZIhvcNAQEFBQADggEBAGZn7nEQUpBVu1iNg5lkvl+EWrk03s6LQt/8
wWb3cKVMkWcO2dzeETukzUoAdJTBrhH+jg5Hvb53jnXzs3j1uNvfppPH5PFNxTYU
YH8v6SmALBoFJXVTvZAqYRnKsTfBgatkbpEDMKjoq6JnOIKU8YMfnByBDfID9jUL
r7Qf2YDuK9YwnZThKb8RFIX3KPnQpxBbdUQeC6jXZD2IU6Id0TFFX8sqBOJkDTZK
FoqCxbUl6DWRY1hiZcM1qMTHJmy2vp13BDjNk5qeVB6QgpRFKG7C7zJ/PHjs8E13
/0MIQerhu36y6wh/2UIdDEY1Ga1Wb7bbX/LYcKybY6BosF14wxM=
-----END CERTIFICATE-----
`

const testA = `-----BEGIN CERTIFICATE-----
MIIDUjCCAjqgAwIBAgIJANYuoGKAZHS+MA0GCSqGSIb3DQEBBQUAMD0xCzAJBgNV
BAYTAlVLMRAwDgYDVQQIDAdFbmdsYW5kMRAwDgYDVQQKDAdFeGFtcGxlMQowCAYD
VQQDDAFCMB4XDTE2MDMwMTE1MjIwMVoXDTE3MDMwMTE1MjIwMVowPTELMAkGA1UE
BhMCVUsxEDAOBgNVBAgMB0VuZ2xhbmQxEDAOBgNVBAoMB0V4YW1wbGUxCjAIBgNV
BAMMAUEwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDP8EKj+C3KlGhw
+e3s6P239N0gW96FPbAYeoyV8Hbz3A97MI+mAeRmhS85aX1Tg3jJwX3J+lLDtCeC
jUQZ4IJUtR99JZaBni7rXB7zFgW6ppoIV1Oy0dMfgm4YVdyv3ww2335BsENtE7P1
p6g9qCa9RItYMLQVrV2TeFqTyPCyVaGfmzhCrBKcIXenBQGSz4SoI6sZgyn6t7n5
Uav4HvU0dd3jqXYuWA+TIaxokRva5EB/K+PsUKtSxYQILAlcpwwwWARvinfJv0Np
W7gBtjOX0CQgTpY91JBrGpxpEa0J5qRnrudtRK8FEJVic9Sn35b+w2j/0T2KbSXu
GfcA5z3PAgMBAAGjVTBTMAwGA1UdEwQFMAMBAf8wOAYIKwYBBQUHAQEELDAqMCgG
CCsGAQUFBzAChhxodHRwOi8vd3d3LmV4YW1wbGUuY29tL2IuY3J0MAkGA1UdEQQC
MAAwDQYJKoZIhvcNAQEFBQADggEBAHjwlz+C9Qn+Ggc1G7TcoIOuA4/yD8KqueIF
GvUYXGgyXtV4cTQ5yTppWy8yhR2ZOCU7llOX0aoS3Oo3fKN7tcQGqz4n5LoPir4z
1A/h8aplp/Fd6xyNdIcvjCH0lvbSgXr/ZwC+Y5uTBZ4q9mYa3VfyQwvf4WnLEYCV
vxsEcab0f5Z9As8rEFb44Dgn5Qj9TMbJ5OqkGocX/fEe+fgSWzgYQQsgQ62r8EOY
QeqTQVAp7z9430uwhPAVKFWT7gQcF9+zAtA93Zwvc3L5b84iwazEpy95EeyCMWzb
4ymNu6ExYn88Tin40xYDjaX7mswccu3drC2icN0AfYEYl0rcD9Q=
-----END CERTIFICATE-----
`
