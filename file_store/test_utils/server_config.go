package test_utils

const SERVER_CONFIG = `
version:
  name: velociraptor
  version: 0.4.6
  commit: 1edf062
  build_time: "2020-07-12T22:09:46+10:00"
Client:
  server_urls:
  - https://localhost:8000/
  ca_certificate: |
    -----BEGIN CERTIFICATE-----
    MIIDTDCCAjSgAwIBAgIRAJH2OrT69FpC7IT3ZeZLmXgwDQYJKoZIhvcNAQELBQAw
    GjEYMBYGA1UEChMPVmVsb2NpcmFwdG9yIENBMB4XDTIxMDQxMzEwNDY1MVoXDTMx
    MDQxMTEwNDY1MVowGjEYMBYGA1UEChMPVmVsb2NpcmFwdG9yIENBMIIBIjANBgkq
    hkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsLO3/Kq7RAwEhHrbsprrvCsE1rpOMQ6Q
    rJHM+0zZbxXchhrYEvi7W+Wae35ptAJehICmbIHwRhgCF2HSkTvNdVzSL9bUQT3Q
    XANxxXNrMW0grOJwQjFYBl8Bo+nv1CcJN7IF2vWcFpagfVHX2dPysfCwzzYX+Ai6
    OK5MqWwk22TJ5NWtUkH7+bMyS+hQbocr/BwKNWGdRlP/+BuUo6N99bVSXqw3gkz8
    FLYHVAKD2K4KaMlgfQtpgYeLKsebjUtKEub9LzJSgEdEFm2bG76LZPbKSGqBLwbv
    x+bJcn23vb4VJrWtbtB0GMxB1bHLTkWgD6PV6ejArClJPvDc9rDrOwIDAQABo4GM
    MIGJMA4GA1UdDwEB/wQEAwICpDAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUH
    AwIwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQUO2IRSDwqgkZt5pkXdScs5Bjo
    ULEwKAYDVR0RBCEwH4IdVmVsb2NpcmFwdG9yX2NhLnZlbG9jaWRleC5jb20wDQYJ
    KoZIhvcNAQELBQADggEBABRNDOPkGRp/ScFyS+SUY2etd1xLPXbX6R9zxy5AEIp7
    xEVSBcVnzGWH8Dqm2e4/3ZiV+IS5blrSQCfULwcBcaiiReyWXONRgnOMXKm/1omX
    aP7YUyRKIY+wASKUf4vbi+R1zTpXF4gtFcGDKcsK4uQP84ZtLKHw1qFSQxI7Ptfa
    WEhay5yjJwZoyiZh2JCdzUnuDkx2s9SoKi+CL80zRa2rqwYbr0HMepFZ0t83fIzt
    zNezVulkexf3I4keCaKkoT6nPqGd7SDOLhOQauesz7ECyr4m0yL4EekAsMceUvGi
    xdg66BlldhWSiEBcYmoNn5kmWNhV0AleVItxQkuWwbI=
    -----END CERTIFICATE-----
  nonce: rKNKAYam310=
  writeback_darwin: /etc/velociraptor.writeback.yaml
  writeback_linux: /tmp/1/velociraptor.writeback.yaml
  writeback_windows: $ProgramFiles\Velociraptor\velociraptor.writeback.yaml
  max_poll: 600
  windows_installer:
    service_name: Velociraptor
    install_path: $ProgramFiles\Velociraptor\Velociraptor.exe
    service_description: Velociraptor service
  darwin_installer:
    service_name: com.velocidex.velociraptor
    install_path: /usr/local/sbin/velociraptor
  version:
    name: velociraptor
    version: 0.4.6
    commit: 1edf062
    build_time: "2020-07-12T22:09:46+10:00"
  pinned_server_name: VelociraptorServer
  max_upload_size: 5242880
  local_buffer:
    memory_size: 52428800
    disk_size: 1073741824
    filename_linux: /var/tmp/Velociraptor_Buffer.bin
    filename_windows: $TEMP/Velociraptor_Buffer.bin
    filename_darwin: /var/tmp/Velociraptor_Buffer.bin
API:
  bind_address: 127.0.0.1
  bind_port: 8001
  bind_scheme: tcp
  pinned_gw_name: GRPC_GW
GUI:
  bind_address: 127.0.0.1
  bind_port: 8889
  gw_certificate: |
    -----BEGIN CERTIFICATE-----
    MIIDQTCCAimgAwIBAgIQWOXRF+ijM59sWOG4CtB+jjANBgkqhkiG9w0BAQsFADAa
    MRgwFgYDVQQKEw9WZWxvY2lyYXB0b3IgQ0EwHhcNMjEwNDEzMTA0NjUyWhcNMjIw
    NDEzMTA0NjUyWjApMRUwEwYDVQQKEwxWZWxvY2lyYXB0b3IxEDAOBgNVBAMMB0dS
    UENfR1cwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQC1Ef/kfXJrdWSw
    Zvm7jEDxyfxi7WoUpV02d4nj2dNbvJyPkFDXzjgzGZR4zie/qFE4kjnaTpQcslau
    +uLbqCFG93t0LDxHEx4OhDz4zB6zSz/0GTYFiLI7jauQvjj+s/IEdMpN1SKLeXuz
    15XiZP3fK2s/N7/+YsEE23y+9RY3qy86xjO5YRv1koPAiJa70uY2s1nUkfDhNrZP
    o+huL4gkiQz2ks2t3HDPXpvyem1lxggiqcwXyEe3ipUK5EbRbVmt19rq2fBd5KVA
    4i9KdY877GWyfvUxDYJKInJE8CnR+gRVHtLLBaxroL0qFWanPpc9LiUw5KCFPkcz
    xh/PmxfjAgMBAAGjdDByMA4GA1UdDwEB/wQEAwIFoDAdBgNVHSUEFjAUBggrBgEF
    BQcDAQYIKwYBBQUHAwIwDAYDVR0TAQH/BAIwADAfBgNVHSMEGDAWgBQ7YhFIPCqC
    Rm3mmRd1JyzkGOhQsTASBgNVHREECzAJggdHUlBDX0dXMA0GCSqGSIb3DQEBCwUA
    A4IBAQCbNYLMGS3ldxn/yDXq+uFLEq29imVSdhUzrBAGXVfH9pa//L63rTjqFHHd
    aTLIAmZS4nwE4Grd2GeQ5ibP3NubtR3XR9QMgfPGJIok/oQ2TQKFqmUhNXgqgMAo
    yq6nMCLNFJnB5UkMmifaDIvHEadDSpbh7RNQlhiZGfWzHrMGl/jGJOqsGO7wjsud
    f2KOTbpIJXkPW8iY50tQuHn6dBjU5cUkh2CDRdjBqhEIT+IgclncJW+HMYznwf6U
    OhDWkHZBDH2OQldv/8JpapmAd/vWPmf7o6KBhFOE7hMW+YiLM7FlOp9sOlSoF6TN
    00DS+oXzYLINPteqVVZRHFowLflT
    -----END CERTIFICATE-----
  gw_private_key: |
    -----BEGIN RSA PRIVATE KEY-----
    MIIEogIBAAKCAQEAtRH/5H1ya3VksGb5u4xA8cn8Yu1qFKVdNneJ49nTW7ycj5BQ
    1844MxmUeM4nv6hROJI52k6UHLJWrvri26ghRvd7dCw8RxMeDoQ8+Mwes0s/9Bk2
    BYiyO42rkL44/rPyBHTKTdUii3l7s9eV4mT93ytrPze//mLBBNt8vvUWN6svOsYz
    uWEb9ZKDwIiWu9LmNrNZ1JHw4Ta2T6Pobi+IJIkM9pLNrdxwz16b8nptZcYIIqnM
    F8hHt4qVCuRG0W1Zrdfa6tnwXeSlQOIvSnWPO+xlsn71MQ2CSiJyRPAp0foEVR7S
    ywWsa6C9KhVmpz6XPS4lMOSghT5HM8Yfz5sX4wIDAQABAoIBAEkOTBqaarHkmY16
    3sPG4OUtQ8F+XNCbr7IsJNxjpJ1cyiY6agPIVnB3n6nmM69mb+7NQL5N+gBiDAtp
    emJ5wYNqksrali9vDS96N0eCw9g/Qjxwd5VgA9l3XYN23HQXkkwt+vnQRrHTgA+d
    qgHAEZYbZJgLHFG88nfNPtWk+SPYJcD+dZUiLJI3OC9S0PdA7Zy1smN1X5nVouQ+
    IDsxnRirsWSm5sXKKC3ZtAsmvMUxmzCvmuS1ctsjdcXSg+NAgUn/GT26SO84hAgx
    MjpkWNHx6ILz+/pdTBIzdGAy3FGu0+nt6+ty2Cyy3sGhYRTkAUUtMvl+9aecIobj
    tbbexgECgYEAznpanIF8Mxe3LHh8k+T5BH4kTvX2cFYHYEHIcFxnVI+0M+A3bex8
    c5Z9o7ScDX1QtFK+dmkRaC2NhAVg1IXVmvi98aW5oRaktQkeSoFzmYpaVQCWeL38
    zDi9v1rHC+SlKA4/ftWKbqZrXNoAyIH4t3mYcBlZ+WgzeZXsFqgAi4ECgYEA4H+h
    EOtpguCR1ajLeJECYYpdcV3nKbog51FKDBtyHDUStmyKFV/4KhEvfPyvTboKjit6
    1sYBQE9ubcckW18gWLx+gpROgPMVYqqdjPU/MY/GXFGrJ3Eiz7+RwFYx5xrS4TAQ
    bzCIFSuXNDzLJapiTC+lWfFlmZerxIGKXrb7pWMCgYAM2eKv4QNI7s3wz6L5VYQG
    Oo+p9ncbMM+OtylxjN7FjZgqrmbkyZZJX3F8ce5QjJ8sSgYeupVhxYwMToVjVN0h
    ox+3XZg7JXyk4+dlvb7jLJBIwBudcj28vD9bHJxM6jM5VgGvtDzEfEksY0kvMZ+4
    U7IgMAPAmTKy+5jRJC28gQKBgGWFsgvZubRC5BBQLfuO8Bm2Hb23rmLzVg/ldev1
    uWvCIH04plU2Oh6j+Oq3xnsK590VEo7vqf8wLIseZHWhlhoq27cnZDzyMLiW2Xq2
    KWa7UwPKtvC2O/zg8AOftx+AxN6rArOVqeSQ+ubXmzKkU6GKZmc/QCfuK/JIdo+Q
    hkprAoGAK+PClPTWvRZ1u1/Yqe3s3ycwMlSazz0q683UnmPRF5De82eq3YKxZCE1
    /DvNPmzfT5Il+lz9nggZFReqRLweZAg7JtH4XwdHSTctLKHLWHyHItqPrsIfY6TF
    p6wdtIM8lu4/FithXYBC2WtYm1JQ8OLd+R71KSxK87xsIqKCD7s=
    -----END RSA PRIVATE KEY-----
  internal_cidr:
  - 127.0.0.1/12
  - 192.168.0.0/16
CA:
  private_key: |
    -----BEGIN RSA PRIVATE KEY-----
    MIIEowIBAAKCAQEAsLO3/Kq7RAwEhHrbsprrvCsE1rpOMQ6QrJHM+0zZbxXchhrY
    Evi7W+Wae35ptAJehICmbIHwRhgCF2HSkTvNdVzSL9bUQT3QXANxxXNrMW0grOJw
    QjFYBl8Bo+nv1CcJN7IF2vWcFpagfVHX2dPysfCwzzYX+Ai6OK5MqWwk22TJ5NWt
    UkH7+bMyS+hQbocr/BwKNWGdRlP/+BuUo6N99bVSXqw3gkz8FLYHVAKD2K4KaMlg
    fQtpgYeLKsebjUtKEub9LzJSgEdEFm2bG76LZPbKSGqBLwbvx+bJcn23vb4VJrWt
    btB0GMxB1bHLTkWgD6PV6ejArClJPvDc9rDrOwIDAQABAoIBAAo6vUIBWEn+MBzD
    SAi080S3cNZFftVUNIfpAObjcgr+Rv/0eeHPSHlvd1wC23eyU2p0UC4j75b/OM/F
    t/z0a1aKAxkF5M/KFk/dWy7FGcWIvcWEbl9GoAPuaBfnKR0tDVmOEsy0P08HdU8L
    9+UCYiBvAK1eQlD3oGA7pvB/9DpHKLSiZOBtmss0EXuJdixKvlcF6GPHBpAjG90g
    ogwcRXJt8qJm9/N5pz+3odYFttXwBn7bdxNLBaUkG3RvrFHUslmN7V0tvFIpjAIT
    f7/5jmLhJugoP6wl9hUEsUSrcdRmSYKRNuHFU06OazBTlka4ksM3z2RFJ6TRhxXZ
    s8U8o3ECgYEAwYKeDJQcx+gRC26Vq6EWT5oHZOLrTh5QrZv/cBo0YP8nhLR0uzwz
    HNj8sMgyFV8yLCYvWaqgRCfCwMoMAUQCH5q0GPNxlQuaL+3WjcTwQeTPms9IuMFh
    rTDt1mi3xPwc5n8ZNafB8+1cNJKOCvrKXdxM/kmRIJVUaFREjyM+LgUCgYEA6cOT
    sl2fp80n10VONcFeVIEaN+YjBapDBJzaNThxTVzjBRsPyUzgEIhQ6r6V8LmG56Wo
    VfyELuvNHgKYvA6mIlsH6l3SLq+F7ohwEDVikp0yzjiMRRhhxQUsnahtHhX3JsUd
    yX2hQOLaaNfNV7gYx64a4iWizFrEa9J2wSUQuD8CgYEAmHZD9h8gCfTysPIg5EeX
    34G4/6i1wieqYw58lCNhT2bZCPpw2jBVCQ6BEPu6UhJd4mD3f4sqmGhHTkQib0DY
    93OZH+t2evrYMZkPKUWYEiKn2w4j+sUKIz1gtkRtPbtxPb237AlPi9NgiV9KoKX1
    mTwAQX1O5cAh780s8yXOUM0CgYA/zC6c+Uw/YZBEAhgsN4/lBC8Bnn9kZmlP8vbi
    m3rgoD8c/5u5Vo+4M1vSFR2ayyd0RRPCE96HZ7ddP1wrxtu0eJ+aaOyZ7TFiPj5H
    TiqO1PQur+QoX1Ufjh/1Dyhok5oWLKnKeczuhnsRLgROsmGg7XVMzvS1TPhabOAY
    KmN7xQKBgEnOjlbCT24fvolHxSJETuoq5IHjwnB/DKTMfnsFfqDPgC/rljqQMF5v
    yzPC/h0xqCh/dI7pIsJ5FjEXOtIJT/sWa1iddB7WC2oFh6AIrVJszt0dQx+4lS2m
    OgdvbViAVYsGELhg/EeJs/ig1v27BMcv2aQtZXTEHXmOd2xL93l5
    -----END RSA PRIVATE KEY-----
Frontend:
  hostname: localhost
  bind_address: 0.0.0.0
  bind_port: 8000
  certificate: |
    -----BEGIN CERTIFICATE-----
    MIIDVzCCAj+gAwIBAgIQTykoXOtft5rT1z1zV0fV9DANBgkqhkiG9w0BAQsFADAa
    MRgwFgYDVQQKEw9WZWxvY2lyYXB0b3IgQ0EwHhcNMjEwNDEzMTA0NjUyWhcNMjIw
    NDEzMTA0NjUyWjA0MRUwEwYDVQQKEwxWZWxvY2lyYXB0b3IxGzAZBgNVBAMTElZl
    bG9jaXJhcHRvclNlcnZlcjCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEB
    ALhp6QuNVcnmtDEhiNlt5OpDHhrd2/Ig+Zp4vaeGcSH/PwGdfvFFAztYX8o1YGKz
    gz4gZ0MiHNEMULXl5ID9rz+4AGzrHRTlKe5RtXnf65nkiN58TdR0JCQkrA/lgXGg
    FBJxzeLedQ+ww7xSE6MHFhpyzendK+WlQW2iU0ZR5X/SDuMG7slrNpT6ipszKEeH
    0YlLpF0uQ38c7PvPIzsnfo2gLmOKlvUa9Ep7TgJPj0Gt149wnWNvUskAhMJhj8cu
    9tWtL45sOGPFOXMVdTE3TS5NCpV1bRe3emHYJWnV7GpLkgkfMtAebtzNZefUWGmW
    5w64cwdoqjPsf44NGAldyjECAwEAAaN/MH0wDgYDVR0PAQH/BAQDAgWgMB0GA1Ud
    JQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAMBgNVHRMBAf8EAjAAMB8GA1UdIwQY
    MBaAFDtiEUg8KoJGbeaZF3UnLOQY6FCxMB0GA1UdEQQWMBSCElZlbG9jaXJhcHRv
    clNlcnZlcjANBgkqhkiG9w0BAQsFAAOCAQEAiXelkRKJXBpxYoQ2zIXL3usS2kzA
    QqFOuehAm7DmM51nW3ArufUjSArbeqCGMEfSlDgdNAIaYEOW/H9C/m2DtboHiyiF
    KKC+VnYBbp866tdO8gO7IDP3M8tBtx7iY00ujgIAaK1W7EWJUTjY8+Ttbjseg8qh
    bQe/oe3PtiO8v17zgnS71Xd8H7PTee8jRdwQbNfIP3Lyn33Yk/lcgmzwgQWnzm/k
    iFslO+2qDS7cmM1JcAMvo/DHE5nhy+PeR+nBAQwO10RG3OFdpfjd1H9JRn3XHWH8
    4nqxOW1yZ1r2CIraWaLnUJDodi+C88Qb/bbS27Xj7xlyNRcW5FzOY+5HRQ==
    -----END CERTIFICATE-----
  private_key: |
    -----BEGIN RSA PRIVATE KEY-----
    MIIEpAIBAAKCAQEAuGnpC41Vyea0MSGI2W3k6kMeGt3b8iD5mni9p4ZxIf8/AZ1+
    8UUDO1hfyjVgYrODPiBnQyIc0QxQteXkgP2vP7gAbOsdFOUp7lG1ed/rmeSI3nxN
    1HQkJCSsD+WBcaAUEnHN4t51D7DDvFITowcWGnLN6d0r5aVBbaJTRlHlf9IO4wbu
    yWs2lPqKmzMoR4fRiUukXS5Dfxzs+88jOyd+jaAuY4qW9Rr0SntOAk+PQa3Xj3Cd
    Y29SyQCEwmGPxy721a0vjmw4Y8U5cxV1MTdNLk0KlXVtF7d6YdgladXsakuSCR8y
    0B5u3M1l59RYaZbnDrhzB2iqM+x/jg0YCV3KMQIDAQABAoIBAQCL0QC3mXipkxib
    B2nSL9eXJTv18q+HbCjSTSi9ktjDhdonR1mvYzuICwDVNpLVQnUG7VpdM8fVVxj6
    UEpm4e+0x7TK35c0+GpKHfkRZgwiqXp79JKtuUCNhqBNjRwPIecV9OK8Vx05wDMm
    3LO2rvC7LL+hx+Y4PorUrHWmowD1XunQdPXuQl+nMHOGLkUZlROJmv1ojpODL90O
    1UC2LaoSKhuSId6DcTXrUqn+CsIkIL/DXiasE2dtcqhbfVhbRO+wLP5rUidI0w3i
    1fh/E4W8B1UROTdfuUjqym5p35gnNCkxXzJs7RcrnWNP6uwH2f5MQ1pW2YlDK7yT
    VEm7I2qBAoGBAPSfqYobCoPOrmQFosAjvrYB+ghKJYlwuY7g/Htf83MWjjuEIMly
    SCECP//Dl9jRMLTeR/3c1rMLFwV0BBihT6alrGyYjGXlLMZnLXOolAhQgXp3sQdy
    FPY+dvRbOGnzfjaM6YVS4f/5e8Or/RAV1TEzd17ZGLxGlLH55pEo8Hw1AoGBAMD9
    bPDDBqZHGcTyn47NYhRShvMl5BrdxpN83d5+Gh4tuSMoTGld+g6cjx0xYUd9C4/l
    eVtsaZ5Le3/5ombnuxFtuOLR2Pd+AaRvgiQhU6dgt1HVkEEVH7sPSiUfhks9PaGG
    orgFAgxqnIfo68nRO2SgWKmKaaX6ltNQDiQ0pf2NAoGBAMAEL/2jOj66VrNWpSjz
    JfkSViR5pztaS512x1lEuQFG3ECld2MJHMskS/5ElwXvNq9Sp+oCllkWRGzYY3q9
    7JrgAxd/Kv20xc+7H7dgxUo9f1UVbNVuXDvqTk+C5yZrpk/FjmxD4fie6eJvN4aP
    zLgBMRs3iC4JrSCDC9Q38YVdAoGAMSsf60M02X454YSTBDXtbBIOkbowoGuqYP8q
    I7pq+w9ZIMzuktW/lKNsjtp611TTckjbn2BClHLUtykpqCAX9/vFY/xk08FFt9g1
    BvF6t6Ubb30CsKIqIGVn1YdvlpG2twVvRYT5HaK32KSOFi/5esyjaiZ6fjAW5yV5
    RN4pAC0CgYBihD0dCFaeH2rZ1GnF6dUak2bJDdwltNMjOYkRQtqFSx4+p4kbQioc
    arboLgY+EgcYdJU+yvyI8sMgZ0vX3Gideado2P9EmXqMx3xoPP19mTMrKLTt/oPb
    S4H9WA8izYRrxB2CiDYib5VWRG62cG4b2c9ZSBDgtkOYMaMH/NhSmw==
    -----END RSA PRIVATE KEY-----
  max_upload_size: 10485760
  dyn_dns: {}
  default_client_monitoring_artifacts:
  - Generic.Client.Stats
  expected_clients: 10000
  GRPC_pool_max_size: 100
  GRPC_pool_max_wait: 60
Datastore:
  implementation: Test
Writeback:
  private_key: |
    -----BEGIN RSA PRIVATE KEY-----
    MIIEogIBAAKCAQEArmgftoc6pi/ZMGZO40UIKXlscTXrZWifDtTGsAhXfaKG4xzu
    LLLIM4Cr+L3ctYgFkWyczXst6Tx6zRyU/l2OqaWmJjhNwXlRwNajx+2ZqTa5zA8r
    lr+QeYrg19+Acmgb8DkPwp8in/f3tHl7Na8U8GE/3CX4nMsLOzcfAEdH/4IRh3b0
    3VW361dlBL8Sw2KJ7ECmhujjtlxu7BUDolxxf8bIkFDVt/nhs9xxm2yI+b2xQnsy
    LDHpsZzSuXj/M38s8u0r59QtJ+ByjFjte+gjGpTc9WlMytTvI/RJUbiEKwOPjBVn
    BcV/1IZ08KokSfhq4xpVY/GPZVL4CEf/ZOo9rQIDAQABAoIBAFnNUW75yHAjuRBb
    zYjmVaKNXBIa8l8f9K59TuT7FpmhIxU0I0surzkdqu8ES+3I4R0VMNP49hXfR1fv
    vKQQ5lFh8uBBI4BYiIjjvCdIp1Ni015H/Wi8sJZ0tPtSoN/HzYLuzremmvyFgK0T
    1CY7RWvUlz4y6wVI4zqVUkgha+gaZjoQklzamwqKHQwqtFyPVISmSp6XL/zexk95
    GVUGps4mtsXWUsSnsmlUD5Ola/7hXeEgbD1nj2Znobu9z0y8mDlIhpFpQJwu36KB
    3o3tqBOXuoukxmsvuW8QxW1xzCICuh8CU5g6kWkyNJOsf4X5Y2js/5Zo9dbLkOrd
    VEnnV+ECgYEAydqTmbWQyOeUxV9mfD2BbjnzLvxMCCggW/i4TGYhtLweO7UPSiQT
    /zK0KX317vupUou//vGEcFKLPVu4xchsGrayOVCWEpurqvZfPmg9lWyF2fi8rZK0
    vOWCw8HIgIbb8EvRCH1v0gNMdzjaf1qLN28W5H/7re4rruQOEuyv29kCgYEA3TC9
    XFAVSePV/Ky22AdbccVacABmM5RAneot/E7DTrA9uGujUB+9kCPIDsPLCjT2uXj/
    yP/a210t8KZBtvW+1Ums06titw65lkG7rjapB08vjF1aD0bjPE4R1uapm+CM6dlm
    oc3Beb8kyA+bXZMpnJT1KtAI3/zrdlZkhQlAL/UCgYAs/uViIUAqGL1oFfERhuBg
    Qti7w4/rTY6REet7VFT1Je4TXzQOUeaHP7U7fpGg+UZwWSiuWwYrx6q0Pcr9g8Td
    W5Z1AkrB0SO+U3c9wRzhPzTDNxhQFODnLr4shvj79ZP3h98L5nJTvVqBRRIny3Y3
    IDNZMlJXHj1smfetLkexWQKBgBgcgAfYEvoDBAiPKz9RTf6Q7NLYuEtXFdQg+vJO
    A6xIOfIoiZzqWNeljuFNJozuSRbewcM/YLQY7DEXboJrN2o4pcZNIG2kBUcD01mi
    S7qoPx6l7nNL3ulr+TXb3xFG4RV8xVtN+pEy7OeCDAWfTSHseu030D/aajB0KnD2
    GTEhAoGARB/E6j/WX+CBPWiF4XLV03F1hEMYY/ZSfijcZQniCNtRQUuIkTMiSv1E
    LZ5KmiY35bmYwkGOST6sd9T586nNEdIfs2ngcXwRcgPmQU7VaKQdeVnxhEG2xXFG
    NtyI/STijkpVi99wF39BvXkQGdJuDjAArjGj5kevCpvyveudL5g=
    -----END RSA PRIVATE KEY-----

Mail: {}
Logging: {}
Monitoring:
  bind_address: 127.0.0.1
  bind_port: 8003
api_config: {}
obfuscation_nonce: RzlAlmdcUyw=
`
