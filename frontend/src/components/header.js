import React, { useEffect, useState } from 'react';
import Heading from './Heading.js';

const Header = () => {
  const [currentIndex, setCurrentIndex] = useState(0);
  const [carouselImages, setCarouselImages] = useState([]);

  const fetchCarouselImages = async () => {
    try {
      const response = await fetch(
        'http://localhost:8080/admin/carousel/active'
      );
      if (response.status === 200) {
        const responseData = await response.json();
        console.log('Carousel API response data:', responseData);
        setCarouselImages(responseData.data);
      } else {
        throw new Error('Network response was not ok for carousel images');
      }
    } catch (error) {
      console.error('Error fetching carousel images: ' + error.message);
    }
  };

  useEffect(() => {
    const slider = document.querySelector('.slider');
    // const slides = document.querySelectorAll('.slide');
    let currentIndex = 0;

    function nextSlide() {
      currentIndex = (currentIndex + 1) % carouselImages.length;
      updateSlider();
    }

    function updateSlider() {
      slider.style.transform = `translateX(-${currentIndex * 100}%)`;
    }

    const sliderInterval = setInterval(nextSlide, 3000);

    return () => {
      clearInterval(sliderInterval);
    };
  }, [currentIndex, carouselImages]);

  useEffect(() => {
    fetchCarouselImages();
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
            <div className="slider">
              {carouselImages.map((image, index) => (
                <div
                  className={`slide ${index === currentIndex ? 'active' : ''}`}
                  key={index}
                >
                  {console.log('Image URL:', image.Image)}
                  <img
                    src={`http://localhost:8080/admin/get/image/active/${image.ID}`}
                    alt={`Slide ${index + 1}`}
                  />
                </div>
              ))}
            </div>
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
                style={{ display: 'none', marginRight: '105px' }}
              >
                <li>
                  <a href="#">Home</a>
                </li>
                <li>
                  <a href="#">Portfolio</a>
                </li>
                <li>
                  <a href="#">About</a>
                </li>
                <li>
                  <a href="#">Contact</a>
                </li>
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
