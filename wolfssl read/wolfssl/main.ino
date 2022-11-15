#include <Arduino.h>
#include <U8g2lib.h>

#ifdef U8X8_HAVE_HW_SPI
#include <SPI.h>
#endif
#ifdef U8X8_HAVE_HW_I2C
#include <Wire.h>
#endif
//U8G2_SSD1306_128X64_NONAME_F_HW_I2C u8g2(U8G2_R0, /* reset=*/ U8X8_PIN_NONE, /* clock=*/ 22, /* data=*/ 21);   // ESP32 Thing, HW I2C with pin remapping
U8X8_SSD1306_128X64_NONAME_SW_I2C u8x8(/* clock=*/ 22, /* data=*/ 21, /* reset=*/ U8X8_PIN_NONE);   // OLEDs without Reset of the Display



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
  u8x8.begin();
  u8x8.setPowerSave(0);
 
}


void loop() { 

rfid = readRFID();
  if (rfid != "")
  {
    Serial.println(rfid);
    Serial.println("Authorized access");
    req["id"] = "asset1";
    serializeJson(req, msg);
    invia(msg, readAsset);
 
    if (reply != "")
    {
      Serial.println(reply);
      printdisplay();
    }
    
    rfid = "";
    msg = "";
  }
}


void printdisplay()
{
  char * p;
  p=strchr(reply,'{');
  Serial.println(p);

  
StaticJsonDocument<256> doc;

DeserializationError error = deserializeJson(doc, p, 256);

if (error) {
  Serial.print("deserializeJson() failed: ");
  Serial.println(error.c_str());
  return;
}

const char* ID = doc["ID"]; // "asset1"
const char* temperature = doc["temperature"]; // "23.5"
const char* humidity = doc["humidity"]; // nullptr
const char* owner = doc["owner"]; // "CHECKTEMP"

  
  u8x8.setFont(u8x8_font_chroma48medium8_r);
  u8x8.drawString(0,1,"ID :");
  u8x8.drawString(5,1,ID);
  u8x8.drawString(0,2,"temperature:");
  u8x8.drawString(0,3,temperature);
  u8x8.drawString(0,4,"humidity:");
  u8x8.drawString(0,5,humidity);

}
  
