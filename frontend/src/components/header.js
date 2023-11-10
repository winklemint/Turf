 import React, { useEffect, useState } from 'react';
 import Heading from './Heading.js';
 import './header.css'
 import ImageSlider from './Imageslider.js';
  

const Header = () => {
  const [NavbarData, setNavbarData] = useState([]);
  console.log(NavbarData);
    useEffect(() => {
        fetch('http://localhost:8080/admin/navbar/active')
          .then((response) => response.json())
          .then((data) => setNavbarData(data.data))
          .catch((error) => console.error('Error fetching carousel data:', error));
      }, []);
  
  const menubar = () => {
    const menu = document.getElementById('mainNav');
    const cross = document.getElementById('cross');
    const tongle = document.getElementsByClassName('hamburger');
    menu.style.display = 'block';
    cross.style.display = 'block';
    if (tongle.length > 0) {
      for (let i = 0; i < tongle.length; i++) {
        tongle[i].style.display = 'none';
      }
    }
  };
  const menubarhide = () => {
    const menu = document.getElementById('mainNav');
    const cross = document.getElementById('cross');
    const tongle = document.getElementsByClassName('hamburger');
    menu.style.display = 'none';
    cross.style.display = 'none';
    if (tongle.length > 0) {
      for (let i = 0; i < tongle.length; i++) {
        tongle[i].style.display = 'block';
      }
    }
  };
  return (
    <div className="App">
      <div className=" container-fuild">
        <header className=" header" />
        <div className="row">
          <div className="col-md-12 col-sm-12 col-lg-12 header-sec">
            
              <ImageSlider/>
            
          </div>
          <div className="col-md-12 col-sm-4 col-lg-12">
            <nav>
              <div className="menu-toggle" id="menuToggle">
                <div
                  className="hamburger"
                  onClick={menubar}
                  style={{ display: 'block' }}
                >
                  <span className="line"></span>
                  <span className="line"></span>
                  <span className="line-last"></span>
                </div>
                <div
                  className="cross"
                  id="cross"
                  style={{ display: 'none' }}
                  onClick={menubarhide}
                >
                  <span className="line-cross"></span>
                  <span className="line-cross1"></span>
                </div>
              </div>
              <ul
                className="main-nav"
                id="mainNav"
                style={{ display: 'none', marginRight: '12px' }}
              >
                {NavbarData.map((item,index) => (
                <li>
                  <a href={item.Link}>{item.Name}</a>
                </li>
                ))}
              </ul>
            </nav>
          </div>
          <div className="col-md-12 col-sm-12 col-lg-12">
            <Heading />
          </div>
        </div>
      </div>
    </div>
  );
};
export default Header;


