import React, { useState, useEffect } from 'react';
import { FaWhatsapp } from 'react-icons/fa6';
import { FaFacebookF } from 'react-icons/fa';
import { IoLogoInstagram } from 'react-icons/io5';
import { FiTwitter } from 'react-icons/fi';
import { FaTelegramPlane } from 'react-icons/fa';
import { SlSocialLinkedin } from 'react-icons/sl';

function Footer() {
  const [headingData, setHeadingData] = useState([]);
  const [icons, setIcons] = useState([]);

  useEffect(() => {
    const fetchHeadingData = async () => {
      try {
        const response = await fetch(
          'http://localhost:8080/admin/heading/active'
        );
        if (response.status === 200) {
          const data = await response.json();
          setHeadingData(data.data);
        }
      } catch (error) {
        console.error('Error fetching heading data:', error.message);
      }
    };
    fetchHeadingData();
  }, []);

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

  return (
    <footer className="footer">
      <div className="footer-box">
        <p className="footer-p1">{headingData.Footer}</p>
        {/* <p className="footerp"></p> */}
        <hr style={{ color: '#7a7af2' }} />
        <div className="foot-brand-icon me-2 text-light">{renderIcons()}</div>
      </div>
      <div className="foot-copyright">
        <p>
          Designed by <span style={{ color: '#c2c2c2' }}>walkingdreamz</span> |
          Download <span style={{ color: '#c2c2c2' }}>walkingdreamz</span>
        </p>
      </div>
    </footer>
  );
}

export default Footer;
