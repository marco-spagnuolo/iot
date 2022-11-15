// Libreria DHT
#include "DHT.h"
// Pin digitale di arduino connesso al DHT
#define DHTPIN 15
// tipo del sensore: DHT 11
#define DHTTYPE DHT11

DHT dht(DHTPIN, DHTTYPE);
void init_temp_read()
{
  dht.begin();

}

String read_temp()
{
  String temp = "";
  // Lettura della temperatura in gradi Celsius
  float t = dht.readTemperature();
  // Verifica se le si presenta un errore di lettura (e riprova nuovamente)
  if (isnan(t)) 
  {
    Serial.println(F("Impossibile leggere dal sensore DHT!"));
    return "";
  }
  temp.concat(t);
  return temp;
}

String read_hum()
{
  String hum = "";
  // Lettura della temperatura in gradi Celsius
  float h = dht.readHumidity();
  // Verifica se le si presenta un errore di lettura (e riprova nuovamente)
  if (isnan(h))
  {
    Serial.println(F("Impossibile leggere dal sensore DHT!"));
    return "";
  }
  hum.concat(h);
  return hum;
}
