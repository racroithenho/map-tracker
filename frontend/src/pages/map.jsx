import { MapContainer, TileLayer, Marker, Popup } from "react-leaflet";
import "leaflet/dist/leaflet.css";

function Map() {
  return (
    <MapContainer
      center={[16.0544, 108.2022]}   // Đà Nẵng
      zoom={13}
      style={{ height: "600px", width: "100%" }}
    >
      <TileLayer
        attribution='© OpenStreetMap contributors'
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
      />

      <Marker position={[16.0544, 108.2022]}>
        <Popup>Đà Nẵng</Popup>
      </Marker>
    </MapContainer>
  );
}

export default Map;