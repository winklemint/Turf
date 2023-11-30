import React, { useEffect, useState } from "react";

function Brandicon() {
  const [icons, setIcons] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8080/admin/icon/active")
      .then((response) => response.json())
      .then((data) => setIcons(data.data))
      .catch((error) => console.error("Error fetching icon data:", error));
  }, []);

  const handleIconClick = (iconLink) => {
    window.open(iconLink);
  };

  return (
    <div>
      {icons.map((icon) => (
        <i
          key={icon.ID}
          className={`fab fa-${icon.Name}`}
          style={{ fontSize: "45px", cursor: "pointer" }}
          onClick={() => handleIconClick(icon.Link)}
        ></i>
      ))}
    </div>
  );
}

export default Brandicon;
