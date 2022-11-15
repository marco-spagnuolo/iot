#include <ArduinoJson.h>
DynamicJsonDocument req(1024);

int input_s = 0;
//char initLedg[] = "POST /init HTTP/1.1\nHost: 192.168.1.70:8081\nConnection: close\nContent-Length: ";
char updateAsset[] = "POST /update HTTP/1.1\nHost: 192.168.1.70:8081\nConnection: close\nContent-Length: ";
char readAsset[] = "POST /read HTTP/1.1\nHost: 192.168.1.70:8081\nConnection: close\nContent-Length: ";
String rfid = "";
String msg = "";



void setup() {
  initWolf();
  delay(500);
  init_rfid();
  delay(500);
  init_temp_read();

}


void loop() { 

rfid = readRFID();
  if (rfid == "DCD7EE8C")
  {
    Serial.println(rfid);
    Serial.println("Authorized access");
    req["id"] = "asset1";
    req["temperature"]   = read_temp();
    req["humidity"] = read_hum();
    req["owner"] = rfid;
    serializeJson(req, msg);
    //Serial.println(msg);

    //updateAsset aggiorna l'asset presente in blockchain passando tutti i parametri. 
    //L'asset verrà modificato solo se già presente in blockchain  
    invia(msg, updateAsset);
    if (reply != "")
    {
      Serial.println(reply);
    }
    rfid = "";
    msg = "";
  }
  else
  if(rfid != "")
  {
    Serial.println(rfid);
    Serial.println("Unauthorized access");
  }
  
  
  if (Serial.available() > 0) {
    input_s = Serial.parseInt();
  }
  
  if (input_s==2) {
    req["id"] = "asset1";
    serializeJson(req, msg);
    invia(msg, readAsset);
 
    if (reply != "")
    {
      Serial.println(reply);
    }
  }



}
