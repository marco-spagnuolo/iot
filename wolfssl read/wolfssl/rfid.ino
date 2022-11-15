#include <MFRC522.h>  // MFRC522 RFID module library.
#include <SPI.h>      // SPI device communication library.
#include <EEPROM.h>   // EEPROM (memory) library.

#define pinRST 27      // Defines pins for RST, SS conncetions respectively.
#define pinSS 5

byte readCard[4];     // Array that will hold UID of the RFID card read
String content= "";

int successRead;

MFRC522 mfrc522(pinSS, pinRST);   // Creates MFRC522 instance.
MFRC522::MIFARE_Key key;          // Creates MIFARE key instance.

void init_rfid() {
  // put your setup code here, to run once:
  //Serial.begin(9600); // Starts the serial connection at 9600 baud rate.
  SPI.begin();        // Initiates SPI connection between RFID module and Arduino.
  mfrc522.PCD_Init(); // Initiates MFRC522 RFID module.
  //Serial.println("RFID reading process initiated.");    // Prints user commands.
}

String readRFID() // Function that will read and print the RFID cards UID.
{
  if ( ! mfrc522.PICC_IsNewCardPresent())  // If statement that looks for new cards.
  {
    return "";
  }

  if ( ! mfrc522.PICC_ReadCardSerial())    // If statement that selects one of the cards.
  {
    return "";
  }
  content= "";
  //Serial.print("UID: ");    
  for (int i = 0; i < mfrc522.uid.size; i++) {  
    readCard[i] = mfrc522.uid.uidByte[i];   // Reads RFID cards UID.
    //content.concat(String(mfrc522.uid.uidByte[i] < 0x10 ? " 0" : " "));
    content.concat(String(mfrc522.uid.uidByte[i], HEX));
    
  }
  content.toUpperCase();
  //Serial.println("il content");
  //Serial.println(content);

  mfrc522.PICC_HaltA();     // Stops the reading process.
  return content;

}


  
