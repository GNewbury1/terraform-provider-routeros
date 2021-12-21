import routeros_api
import os

def main():
    user = os.environ.get("ROS_USERNAME")
    pswd = os.environ.get("ROS_PASSWORD")
    ip_addr = os.environ.get("ROS_IP_ADDRESS")

    connection = routeros_api.RouterOsApiPool(ip_addr, username=user, password=pswd, port=8728, plaintext_login=True)
    api = connection.get_api()

    certificate = api.get_resource("/certificate")
    certificate.add(name="root-cert", common_name="MyRouter", days_valid="3650", key_usage="key-cert-sign,crl-sign")
    certificate.add(name="https-cert", common_name="MyRouter", days_valid="3650")
    certs = certificate.get()
    root_cert_id = [x['id'] for x in certs if x['name'] == "root-cert"][0]
    http_cert_id = [x['id'] for x in certs if x['name'] == "https-cert"][0]
    api.get_binary_resource("/").call("certificate/sign", {"id": bytes(root_cert_id, "utf-8")})
    api.get_binary_resource("/").call("certificate/sign", {"id": bytes(http_cert_id, "utf-8"), "ca": b"root-cert"})
    ip_service = api.get_resource("/ip/service")
    services = ip_service.get()
    www_ssl_service_id = [x['id'] for x in services if x['name'] == 'www-ssl'][0]
    ip_service.set(id=www_ssl_service_id, certificate="https-cert", disabled="false")


if __name__ == "__main__":
    main()