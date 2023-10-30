import React, { useEffect, useState } from 'react';
import './Home.css';
import { Link } from 'react-router-dom';
import Testimonial from '../TestimonialComponent/Testimonial';
import Swiper from '../SwiperComponent/Swiper';

const Home = () => {
  
  const [data, setData] = useState({
    heading: '',
    subheading: '',
    button: '',
  });

  const [carouselImages, setCarouselImages] = useState([]);


  const fetchData = async () => {
    try {
      const response = await fetch(`http://localhost:8080/admin/content/active`);
      if (response.status === 200) {
        const responseData = await response.json();
        console.log('API response data:', responseData);
        setData({
          heading: responseData.data.Heading,
          subheading: responseData.data.SubHeading,
          button: responseData.data.Button,
        });
      } else {
        throw new Error('Network response was not ok');
      }
    } catch (error) {
      console.error('Error fetching content: ' + error.message);
    }
  };
  
  const fetchCarouselImages = async () => {
    try {
      const response = await fetch('http://localhost:8080/admin/carousel/active');
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
    // Initialize the slider and start it when data or carousel images change
    const slider = document.querySelector('.slider');
    const slides = document.querySelectorAll('.slide');
    let currentIndex = 0;

    function nextSlide() {
      currentIndex = (currentIndex + 1) % slides.length;
      updateSlider();
    }

    function updateSlider() {
      slider.style.transform = `translateX(-${currentIndex * 100}%)`;
    }

    // Only start the slider once when the component mounts
    const interval = setInterval(nextSlide, 2000);

    // Clear the interval when the component unmounts or when dependencies change
    return () => {
      clearInterval(interval);
    };
  }, [carouselImages]);


  useEffect(() => {
    fetchData();
    fetchCarouselImages();
  }, []);
  
  return (
    <>
     <header className="header">
     <div className="row">
          <div className="col-md-12 col-sm-12 col-lg-12 header-sec">
          <div className="slider">
          {carouselImages.slice(0, 10).map((image, index) => (
                <div className={`slide ${index === 0 ? 'active' : ''}`} key={index}>
                  {console.log('Image URL:', image.Image)}
                  <img src={`http://localhost:8080/admin/get/image/active/${image.ID}`} alt={`Slide ${index + 1}`} />
                </div>
              ))}
            </div>
          </div>

          <div className="col-md-12 col-sm-4 col-lg-12">
            <nav>
              <div className="menu-toggle" id="menuToggle">
                <div className="hamburger">
                  <span className="line"></span>
                  <span className="line"></span>
                  <span className="line-last"></span>
                </div>
                <div className="cross" style={{ display: "none" }}>
                  <span className="line-cross"></span>
                  <span className="line-cross1"></span>
                </div>
              </div>
              <ul className="main-nav" id="mainNav" style={{ display: "none" }}>
                <li>
                  <Link to="/">Home</Link>
                </li>
                <li>
                  <Link to="/booking">Booking</Link>
                </li>
                <li>
                  <Link to="/about">About</Link>
                </li>
                <li>
                  <Link to="/contact">Contact</Link>
                </li>
              </ul>
            </nav>
          </div>
        </div>
        <div className="col-md-12 col-sm-12 col-lg-12">
          <div className="text-box">
            <p className="text-p1">{data.heading}</p>
            <h3 className="text-h3">{data.subheading}</h3>
            <p className="text-p2"></p>
            <Link to="/booking">
              <button className="text-button">
                <span className="text-btn-linkk">{data.button}</span>
              </button>
            </Link>
            <div className="icon-sec">
              <p>Join me here</p>
              <div className="brand-icon">
                <i className='fab fa-whatsapp' style={{ fontSize: "30px" }}></i>
                <i className='fab fa-facebook' style={{ fontSize: "30px" }}></i>
                <i className='fab fa-twitter' style={{ fontSize: "30px" }}></i>
                <i className='fab fa-linkedin' style={{ fontSize: "30px" }}></i>
              </div>
            </div>
          </div>
        </div>
      </header>
      <Swiper />
      <Testimonial />
    </>
  );
};

export default Home;

