#include <wolfssl.h>
#include <wolfssl/ssl.h>
#include <WiFi.h>
char reply[1024];

//Connecting to wifi

const char* ssid = "OnePlus 8T";
const char* password = "lello555";



const unsigned char cert[] = \
                             "-----BEGIN CERTIFICATE-----\n" \
                             "MIIDozCCAougAwIBAgIUCNT3YJl4uSQs6rSqgmc/hgLxhxowDQYJKoZIhvcNAQEL\n" \
                             "BQAwHjEcMBoGA1UEAwwTU29ja2V0VG9vbHMgVGVzdCBDQTAeFw0yMTEyMDYxNDA5\n" \
                             "NTBaFw0yMjEyMDYxNDA5NTBaMGMxCzAJBgNVBAYTAklUMQ8wDQYDVQQIDAZJdGFs\n" \
                             "aWExEDAOBgNVBAcMB1NhbGVybm8xHTAbBgNVBAoMFENhdGFseXN0IERldmVsb3Bt\n" \
                             "ZW50MRIwEAYDVQQDDAlsb2NhbGhvc3QwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw\n" \
                             "ggEKAoIBAQCUfTdBEEaFHuk3Fu/qi6dR79y9ZfIc0bOjdWidZFMQxas+5HXpxhjU\n" \
                             "/tIPh9bWn683Gq22jDqUyzX/vRmB/3fCiudXjVbzC8FsxVBKSXZc+eO1dkMIkX9Y\n" \
                             "uzVoksSLjtBKXwlwLZUF8eaRnqLnFF1GBG1SS+YuawPYEDtM2ysW5AGVyfNFbrsF\n" \
                             "x+OirmwoUE9kLyJ1RKqfwEJ4OtTI8NMpVgB51eW4RyR6dkzcvmN9zKVyXFjqFNZ/\n" \
                             "2F9AOkyNg8nXDJFZjqWamFzHWLlpD5lqEmpyPYxROCdfitUozzwIf6kUG1PGb5/W\n" \
                             "PoUD7un/N+rG4GXaSDmHoPVl5+Rcn4O1AgMBAAGjgZMwgZAwHQYDVR0OBBYEFCnv\n" \
                             "AbLXlObIc3nfsf8Uw8N2I/WBMAwGA1UdEwEB/wQCMAAwEQYJYIZIAYb4QgEBBAQD\n" \
                             "AgZAMAsGA1UdDwQEAwIF4DATBgNVHSUEDDAKBggrBgEFBQcDATAsBgNVHREEJTAj\n" \
                             "gglsb2NhbGhvc3SHBKsTn6GHEAAAAAAAAAAAAAAAAAAAAAEwDQYJKoZIhvcNAQEL\n" \
                             "BQADggEBADHLb2xtrb3W9OXYlCT40gbz9gAizOWv1DhTmuHa4LyOxj7ITCn3P93O\n" \
                             "jxOkrGy+MK4slN0AzsS5ZMFaSLIF893nOFe2PVCVLzNpYJS6xyp7LqKFX3dwWx9u\n" \
                             "Ari1V1OHZBZdRNhwQi+MfJkZlZJHMt168vYTouEftLOY/UgT+oB27KfZcVEr7H+U\n" \
                             "TVWvmiVWphn7Dj9N7sywop6HWfBc/HLCEFXot6uGMtlSY7ohOLnxLZpcukP5ipf6\n" \
                             "vuPoIZlVjweiysZ/iIi/B5JheMS/Tk+ycbF/KeE/6PzA6nQyq0vK/3XTvNw/n8QD\n" \
                             "gp7kAIRnrcVEdTeweD9zuRe3BpenDWY=\n" \
                             "-----END CERTIFICATE-----\n" \
                             "-----BEGIN CERTIFICATE-----\n" \
                             "MIIDbDCCAlSgAwIBAgIUccC2YH8D254mdqCgdEMc+iXMyLQwDQYJKoZIhvcNAQEL\n" \
                             "BQAwHjEcMBoGA1UEAwwTU29ja2V0VG9vbHMgVGVzdCBDQTAeFw0yMTEyMDYxNDA0\n" \
                             "MTBaFw0yMjEyMDYxNDA0MTBaMB4xHDAaBgNVBAMME1NvY2tldFRvb2xzIFRlc3Qg\n" \
                             "Q0EwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQC0qf0VLtLFF4N18hGc\n" \
                             "adwip17y3Vr0eh1dPUDgKHLG6HFp8rdkaC+URnaN+x7vYrGXKDpple4YRWB8yjZT\n" \
                             "mmCUOgfNfuXWKL4JfUEn/rxWz9xF4Nz+tNQPbi3rbgeLzW8azI2L+9dufiSXB8ds\n" \
                             "p9bnM90c4wR5drB4qs3WKAr/RFOkDRw4JWBn6tlKHys69mUZABQLoLKb0VBJF4wl\n" \
                             "hpEfso8c+ez46mHjdqUk5uDrKnnDdyyPV0DtEwrnycU2Edn4AzFMHKy0yWGF9Pvv\n" \
                             "AvzZPuicL0WsrGdl92+x0i5DdIfWTx/xcc7avlCsbjCpm0bGn/amh31wI4BZiMfS\n" \
                             "13VDAgMBAAGjgaEwgZ4wHQYDVR0OBBYEFCkRoGz10LZnMNFNSsQVOJKLeQ6tMFkG\n" \
                             "A1UdIwRSMFCAFCkRoGz10LZnMNFNSsQVOJKLeQ6toSKkIDAeMRwwGgYDVQQDDBNT\n" \
                             "b2NrZXRUb29scyBUZXN0IENBghRxwLZgfwPbniZ2oKB0Qxz6JczItDASBgNVHRMB\n" \
                             "Af8ECDAGAQH/AgEAMA4GA1UdDwEB/wQEAwIBhjANBgkqhkiG9w0BAQsFAAOCAQEA\n" \
                             "Hlzc2ZP9rKRCyGSRvRyoQtSBbPA2bZcWi/ZVIdpdWqc2ugcWky9gBZqg4AkHQMDJ\n" \
                             "0FY4x/Ps8vTrOlhd08iik6RGJr0HSkDEAcFNG+AjgsylxXlyK0oekCU2lO/kP42P\n" \
                             "AnF+NghjIkE6PuNj/14a2CorL2sf12pZhoHxdHkoDBkQlCiez36EkkjA+wdLmKXC\n" \
                             "S45zR0KPKUuTj1zSLYMaIo2MHvuBTqpf49GSc9MXLZqR0M2dqmgV2HFQ6Us/4uA0\n" \
                             "IazxhLA2Fg4RxYFywB2PBGpO7of5asMe6f1Is38bpftqn6GBv6hMKrGi2cHEL+2S\n" \
                             "GCor1YL6KVhn2jUnWxiKXw==\n" \
                             "-----END CERTIFICATE-----\n";


//Server address  172.19.163.78
const char host[] = "172.19.163.78";
const int port = 8084; /* port on server to connect to */

int EthernetSend(WOLFSSL* ssl, char* msg, int sz, void* ctx);
int EthernetReceive(WOLFSSL* ssl, char* reply, int sz, void* ctx);
int reconnect = 10;

WiFiClient client;

WOLFSSL_CTX* ctx = NULL;
WOLFSSL* ssl = NULL;


void initWolf() {
  Serial.begin(9600);
  WiFi.begin(ssid, password);

  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }

  Serial.println("");
  Serial.println("WiFi connected");
  Serial.println("IP address: ");
  Serial.println(WiFi.localIP());

  WOLFSSL_METHOD* method;

  method = wolfTLSv1_3_client_method();
  if (method == NULL) {
    Serial.println("unable to get method");
    return;
  }
  ctx = wolfSSL_CTX_new(method);
  if (ctx == NULL) {
    Serial.println("unable to get ctx");
    return;
  }
  /* initialize wolfSSL using callback functions */


  int ret = wolfSSL_CTX_use_certificate_chain_buffer(ctx, cert, sizeof(cert));
  if (ret != SSL_SUCCESS) {
    Serial.println("Errore certificato");
  }

  wolfSSL_CTX_load_verify_buffer(ctx, cert, sizeof(cert), SSL_FILETYPE_PEM);
  wolfSSL_SetIOSend(ctx, EthernetSend);
  wolfSSL_SetIORecv(ctx, EthernetReceive);

  return;
}

int EthernetSend(WOLFSSL* ssl, char* msg, int sz, void* ctx) {
  int sent = 0;
  sent = client.write((byte*)msg, sz);
  return sent;
}

int EthernetReceive(WOLFSSL* ssl, char* reply, int sz, void* ctx) {
  int ret = 0;
  while (client.available() > 0 && ret < sz) {
    reply[ret++] = client.read();
  }

  return ret;
}

void invia(String mex, char header[])
{
  const char *messaggio = "";
  messaggio = mex.c_str();

  int err            = 0;
  int input          = 0;
  int total_input    = 0;



  int msgSz = strlen(messaggio);
  char msgSz_char[5];
  itoa (msgSz, msgSz_char, 10);
  char msg[1024] = {""};
  strcat(msg,   header);
  strcat(msg,   msgSz_char);
  strcat(msg,   "\n\n");
  strcat(msg,   messaggio);

  int post_len = strlen(messaggio);
  char errBuf[80];
  const char* cipherName;


  reconnect--;

  if (client.connect(host, port)) {

    Serial.print("Connected to ");
    Serial.println(host);

    ssl = wolfSSL_new(ctx);
    if (ssl == NULL) {
      Serial.println("Unable to allocate SSL object");
      return ;
    }

    err = wolfSSL_connect(ssl);

    if (err != WOLFSSL_SUCCESS) {
      err = wolfSSL_get_error(ssl, 0);
      Serial.println(err);
      wolfSSL_ERR_error_string(err, errBuf);
      Serial.print("TLS Connect Error: ");
      Serial.println(errBuf);
    }

    Serial.print("SSL version is ");
    Serial.println(wolfSSL_get_version(ssl));

    cipherName = wolfSSL_get_cipher(ssl);
    Serial.print("SSL cipher suite is ");
    Serial.println(cipherName);

    if ((wolfSSL_write(ssl, msg, 1024)) == 1024) {
      Serial.print("Invio: ");
      //Serial.println(msg);
      //Serial.print("Server response: ");
      /* wait for data */
      /*
        while (!client.available()) {
        Serial.println("Sto aspettando il server");
        delay(1000);
        }*/
      /* read data */


      //Serial.println("AO");
      input = wolfSSL_read(ssl, reply, 1024);
      total_input += input;
      //Serial.println("l'input in wolf è" + input);
      if (input < 0) {
        err = wolfSSL_get_error(ssl, 0);
        wolfSSL_ERR_error_string(err, errBuf);
        Serial.print("TLS Read Error: ");
        Serial.println(errBuf);
      } else if (input > 0) {
        reply[input] = '\0';
        //Serial.print(reply);
        //return reply;
      } else {
        Serial.println();
      }


    } else {
      err = wolfSSL_get_error(ssl, 0);
      wolfSSL_ERR_error_string(err, errBuf);
      Serial.print("TLS Write Error: ");
      Serial.println(errBuf);
    }

    delay(1000);
    wolfSSL_shutdown(ssl);
    wolfSSL_free(ssl);

    client.stop();
    Serial.println("\nConnection complete.");
    reconnect = 0;
  } else {
    Serial.println("Trying to reconnect...");
  }

  delay(1000);
}
