import React, { useEffect, useState } from "react";
import Brandicon from "./Brandicon";

function Footer(props) {
	console.log(new Date)
  return (
    <footer className="footer">
      <div className="footer-box">
        <p className="footer-p1">{props.Headingdata.Footer}</p>
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



