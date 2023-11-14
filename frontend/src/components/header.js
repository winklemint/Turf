import React, { useState } from 'react';
import Heading from './Heading.js';
import SliderHeader from './SliderHeader';
import { Container, Nav } from 'react-bootstrap';
import { HiOutlineMenuAlt3 } from 'react-icons/hi';
import { RxCross1 } from 'react-icons/rx';

const Header = () => {
  const [menuVisible, setMenuVisible] = useState(false);
  const [showMenuIcon, setShowMenuIcon] = useState(true);

  const toggleMenu = () => {
    setMenuVisible(!menuVisible);
    setShowMenuIcon(!showMenuIcon);
  };

  const closeMenu = () => {
    setMenuVisible(false);
    setShowMenuIcon(true);
  };

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
                <Nav.Link href="#" className="me-3 fs-3 text-white">
                  Home
                </Nav.Link>
                <Nav.Link href="#" className="me-3 fs-3 text-white">
                  Portfolio
                </Nav.Link>
                <Nav.Link href="#" className="me-3 fs-3 text-white">
                  About
                </Nav.Link>
                <Nav.Link href="#" className="me-3 fs-3 text-white">
                  Contact
                </Nav.Link>
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
    </div>
  );
};

export default Header;
