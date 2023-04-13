import { useState } from "react";

function App() {
  const [kebutuhan, setKebutuhan] = useState({
    gula: 0,
    tepung: 0,
    coklat: 0,
  });
  const [tersedia, setTersedia] = useState({
    gula: 0,
    tepung: 0,
    coklat: 0,
  });
  const [sisa, setSisa] = useState({
    gula: 0,
    tepung: 0,
    coklat: 0,
  });
  const [total, setTotal] = useState(0);
  const hitung = () => {
    const req = new XMLHttpRequest();
    const data = {
      kebutuhan: kebutuhan,
      tersedia: tersedia,
    };
    req.open("POST", "http://localhost:8000/snaki");
    req.onload = (a) => {
      const res = JSON.parse(req.response);
      setTotal(res.total);
      setSisa(res.sisa);
    };
    req.send(JSON.stringify(data));
  };
  return (
    <div className="w-screen h-screen bg-gray-200 pt-32">
      <div className="mx-auto p-2 w-[700px] bg-green-300 flex flex-col">
        <div className="flex justify-between space-x-2">
          <div>
            <div>Kebutuhan per kemasan</div>
            <div className="flex">
              <input
                className="my-1 w-full p-1"
                type="number"
                placeholder="Gula"
                onChange={(e) =>
                  setKebutuhan((prev) => ({
                    ...prev,
                    gula: parseFloat(e.target.value),
                  }))
                }
              ></input>
              <p className="bg-white my-1 p-1">gram</p>
            </div>
            <div className="flex">
              <input
                className="my-1 w-full p-1"
                type="number"
                placeholder="Tepung tapioka"
                onChange={(e) =>
                  setKebutuhan((prev) => ({
                    ...prev,
                    tepung: parseFloat(e.target.value),
                  }))
                }
              ></input>
              <p className="bg-white my-1 p-1">gram</p>
            </div>
            <div className="flex">
              <input
                className="my-1 w-full p-1"
                type="number"
                placeholder="Coklat padat"
                onChange={(e) =>
                  setKebutuhan((prev) => ({
                    ...prev,
                    coklat: parseFloat(e.target.value),
                  }))
                }
              ></input>
              <p className="bg-white my-1 p-1">gram</p>
            </div>
          </div>
          <div>
            <div>Bahan baku yang tersedia</div>
            <div className="flex">
              <input
                className="my-1 w-full p-1"
                type="number"
                placeholder="Gula"
                onChange={(e) =>
                  setTersedia((prev) => ({
                    ...prev,
                    gula: parseFloat(e.target.value),
                  }))
                }
              ></input>
              <p className="bg-white my-1 p-1">kg</p>
            </div>
            <div className="flex">
              <input
                className="my-1 w-full p-1"
                type="number"
                placeholder="Tepung tapioka"
                onChange={(e) =>
                  setTersedia((prev) => ({
                    ...prev,
                    tepung: parseFloat(e.target.value),
                  }))
                }
              ></input>
              <p className="bg-white my-1 p-1">kg</p>
            </div>
            <div className="flex">
              <input
                className="my-1 w-full p-1"
                type="number"
                placeholder="Coklat padat"
                onChange={(e) =>
                  setTersedia((prev) => ({
                    ...prev,
                    coklat: parseFloat(e.target.value),
                  }))
                }
              ></input>
              <p className="bg-white my-1 p-1">kg</p>
            </div>
          </div>
        </div>
        <div className="text-right">
          <button className="bg-blue-400 px-2 rounded" onClick={() => hitung()}>
            Hitung
          </button>
        </div>
        <div>
          <div>Total produksi: {total} kemasan</div>
          <div>
            Sisa bahan baku: Gula {sisa.gula} gram, Tepung tapioka {sisa.tepung}{" "}
            gram, Coklat padat {sisa.coklat} gram
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
