import React, { useEffect, useState } from "react";
import Brandicon from "./Brandicon";

function Footer() {
  const [Heading, setHeading] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8080/admin/heading/active")
      .then((response) => response.json())
      .then((data) => setHeading(data.data))
      .catch((error) => console.error("Error fetching Heading data:", error));
  }, []);

  return (
    <footer className="footer">
      <div className="footer-box">
        <p className="footer-p1">{Heading.Footer}</p>
        <p className="footerp"></p>
		<div className="foot-brand-icon">
        <Brandicon/>
		
       
        </div>
      </div>
      <div className="foot-copyright">
        <p>
          Designed by <span style={{ color: "#c2c2c2" }}>walkingdreamz</span> | Download{" "}
          <span style={{ color: "#c2c2c2" }}>walkingdreamz</span>
        </p>
      </div>
    </footer>
  );
}

export default Footer;



