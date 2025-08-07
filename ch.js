const airportsJSON = require("./data/airports.json");
const { promises } = require("node:fs");

(async () => {
    const fs = promises
   const airportObject = Object.keys(airportsJSON).map((key) => {
    const { iata, city, name, state, country, latitude, longitude } =
      airportsJSON[key];
    return { iata, city, name, state, country, lat: latitude, lon: longitude };
  });
  await fs.writeFile("formatted-airport.json", JSON.stringify(airportObject, null, 2))
})();
