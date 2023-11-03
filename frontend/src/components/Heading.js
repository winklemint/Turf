import { useEffect, useState } from "react";
import { Link } from "react-router-dom";

import React from "react";
import Brandicon from "./Brandicon";
import BookingForm from "./BookNow";

const Heading = () => {
  const [contant, setcontant] = useState({});


  const fetchDataFromAPI = () => {
    fetch("http://localhost:8080/admin/content/active")
      .then((res) => res.json())
      .then((data) => {
        setcontant(data.data);
      });
  };

  useEffect(() => {
    fetchDataFromAPI();
  }, []);

  return (
    <div className="text-box ">
      <p className="text-p1">{contant.Heading}</p>
      <h3 className="text-h3">{contant.SubHeading}</h3>
     <Link to={'/BookNow'}><button className="text-button">
        <a href="#" className="text-btn-linkk">
          {contant.Button}
        </a>
      </button></Link> 
      
      <div className="icon-sec">
        <p>Join me here</p>
        <div className="brand-icon">
          <Brandicon />
        </div>
      </div>
    </div>
  );
};

export default Heading;
