import React, { useState, useEffect } from 'react';
import Heading from './Heading.js';
import SliderHeader from './SliderHeader.js';
import { Container, Nav } from 'react-bootstrap';
import { HiOutlineMenuAlt3 } from 'react-icons/hi';
import { RxCross1 } from 'react-icons/rx';
import Section1 from './section1.js';
import Section2 from './Section2.js';
import Footer from './Footer.js';

const Header = () => {
  const [menuVisible, setMenuVisible] = useState(false);
  const [showMenuIcon, setShowMenuIcon] = useState(true);
  const [navbarItems, setNavbarItems] = useState([]);

  const toggleMenu = () => {
    setMenuVisible(!menuVisible);
    setShowMenuIcon(!showMenuIcon);
  };

  const closeMenu = () => {
    setMenuVisible(false);
    setShowMenuIcon(true);
  };

  const fetchNavbarData = () => {
    fetch('http://localhost:8080/admin/navbar/get')
      .then((res) => res.json())
      .then((data) => {
        setNavbarItems(data.data);
      });
  };

  useEffect(() => {
    fetchNavbarData();
  }, []);

  return (
    <div className="App">
      <Container fluid>
        <header className="header d-flex" />
        <div className="row">
          <div className="col-md-12 col-sm-12 col-lg-12 header-sec p-0">
            <SliderHeader />
          </div>

          <div className="col-md-12 col-sm-4 col-lg-12">
            <Nav className="menu-toggle-nav end-0 top-0 mt-3 position-absolute text-light">
              <div
                className={`main-nav text-white d-flex mt-2 mr-5  ${
                  menuVisible ? '' : 'd-none '
                }`}
              >
                {navbarItems.map((item) => (
                  <Nav.Link
                    key={item.ID}
                    href={item.Link}
                    className="me-3 fs-3 text-white"
                  >
                    {item.Name}
                  </Nav.Link>
                ))}
              </div>
              <div
                className="menu-toggle d-flex justify-content-end"
                id="menuToggle"
              >
                {showMenuIcon ? (
                  <HiOutlineMenuAlt3
                    onClick={toggleMenu}
                    className="display-block text-light fs-1"
                  />
                ) : (
                  <RxCross1 onClick={closeMenu} className="text-light fs-1" />
                )}
              </div>
            </Nav>
          </div>

          <div className="col-md-12 col-sm-12 col-lg-12">
            <Heading />
          </div>
        </div>
      </Container>
      <div>
        <Section1 />
      </div>
      <div>
        <Section2 />
      </div>
      <div>
        <Footer />
      </div>
    </div>
  );
};

export default Header;
