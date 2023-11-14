import React, { useEffect, useState } from 'react';
import { Col, Button } from 'react-bootstrap';

const Heading = () => {
  const [contant, setcontant] = useState({});

  const fetchDataFromAPI = () => {
    fetch('http://localhost:8080/admin/content/active')
      .then((res) => res.json())
      .then((data) => {
        setcontant(data.data);
      });
  };

  useEffect(() => {
    fetchDataFromAPI();
  }, []);

  return (
    <Col className="text-box position-absolute top-0 mt-5 text-white p-5">
      <p className="text-p1 fs-2 fw-bolder">{contant.Heading}</p>
      <h3 className="text-h3 fs-1 mt-3">
        <b>{contant.SubHeading}</b>
      </h3>
      <Button className="text-button">
        <a href="#" className="btn btn-primary fs-5 fw-bolder">
          {contant.Button}
        </a>
      </Button>
      <div className="icon-sec d-flex">
        <p className="fs-3 me-2">Join me here </p>
        <div className="brand-icon fs-3 me-2">
          <i className="fab fa-whatsapp  text-success"></i>
          <i className="fab fa-facebook text-primary"></i>
          <i className="fab fa-twitter text-info"></i>
          <i className="fab fa-linkedin text-primary"></i>
        </div>
      </div>
    </Col>
  );
};

export default Heading;
