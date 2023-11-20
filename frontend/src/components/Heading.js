import React, { useEffect, useState } from 'react';
import { Col, Button } from 'react-bootstrap';
import { FaWhatsapp } from 'react-icons/fa6';
import { FaFacebookF } from 'react-icons/fa';
import { IoLogoInstagram } from 'react-icons/io5';
import { FiTwitter } from 'react-icons/fi';
import { FaTelegramPlane } from 'react-icons/fa';
import { SlSocialLinkedin } from 'react-icons/sl';
import { Link } from 'react-router-dom';

const Heading = () => {
  const [contant, setcontant] = useState({});
  const [icons, setIcons] = useState([]);

  useEffect(() => {
    const fetchIcons = async () => {
      try {
        const response = await fetch('http://localhost:8080/admin/icon/get');
        if (response.status === 200) {
          const data = await response.json();
          setIcons(data.data);
        }
      } catch (error) {
        console.error('Error Fetching Details:', error.message);
      }
    };
    fetchIcons();
  }, []);

  const handleIconClick = (link) => {
    window.open(link, '_blank');
  };

  const renderIcons = () => {
    return icons.map((icon) => {
      const IconComponent = getIconComponent(icon.Name);

      if (IconComponent) {
        return (
          <IconComponent
            key={icon.ID}
            className="text-light fs-3 me-2"
            onClick={() => handleIconClick(icon.Link)}
          />
        );
      }

      return null;
    });
  };

  const getIconComponent = (iconName) => {
    switch (iconName.toLowerCase()) {
      case 'whatsapp':
        return FaWhatsapp;
      case 'facebook':
        return FaFacebookF;
      case 'instagram':
        return IoLogoInstagram;
      case 'twitter':
        return FiTwitter;
      case 'telegram':
        return FaTelegramPlane;
      case 'linkedin':
        return SlSocialLinkedin;
      default:
        return null;
    }
  };

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
        <Link to="/booknow" className="btn btn-primary fs-5 fw-bolder">
          {contant.Button}
        </Link>
      </Button>
      <div className="icon-sec d-flex">
        <p className="fs-3 me-2">Join me here </p>
        <div className="brand-icon fs-3 me-2">{renderIcons()}</div>
      </div>
    </Col>
  );
};

export default Heading;
