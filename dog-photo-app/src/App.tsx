import { useState } from "react";
import "./App.css";
function App() {
  const [imageUrl, setImageUrl] = useState("");

  const breeds = [
    "Australian Shepherd",
    "Dachshund",
    "Pitbull",
    "Golden Retriever",
    "African",
  ];

  function handleClick(breed: string) {
    // Send a request to the proxy server
    fetch(`http://localhost:8080/dog?breed=${encodeURIComponent(breed)}`)
      .then((response) => response.json())
      .then((data) => {
        setImageUrl(data.imageUrl);
        console.log(data.imageUrl);
        console.log(data);
      })
      .catch((error) => {
        console.error("Error fetching image:", error);
      });
  }

  return (
    <div className="w-screen h-screen bg-blue-500 grid grid-cols-2">
      <div className="col-span-1 p-4 items-center justify-center flex flex-col">
        <ul className="space-y-3">
          {breeds.map((breed) => (
            <li key={breed}>
              <button
                className="w-full border p-3"
                onClick={() => handleClick(breed)}
              >
                {breed}
              </button>
            </li>
          ))}
        </ul>
      </div>
      <div className="col-span-1 flex flex-col items-center justify-center p-4">
        <img src={imageUrl} alt="dog" className="max-w-96 max-h-96 border" />
      </div>
    </div>
  );
}

export default App;
