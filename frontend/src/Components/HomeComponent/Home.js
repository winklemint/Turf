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

  const [currentIndex, setCurrentIndex] = useState(0);

  const images = [
    'assets/images/birthdayparties2.jpg',
    'assets/images/DSC00867.jpg',
    'assets/images/sixer-zon.jpg',
  ];

  const updateSlider = () => {
    const nextIndex = (currentIndex + 1) % images.length;
    setCurrentIndex(nextIndex);
  };

  const fetchData = async () => {
    try {
      const response = await fetch(`http://localhost:8080/admin/content/get/1`);
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
      console.log("jhfkjfkjkr")
      console.error('Error fetching content: ' + error.message);
    }
  };
  //By adding console logs, you can inspect the response data and any error messages in your browser's developer console, which can help you identify the issue and resolve it.
  
  
  
  
  
  

  useEffect(() => {
    fetchData();

    const sliderInterval = setInterval(updateSlider, 3000);

    return () => {
      // Clear the interval when the component unmounts
      clearInterval(sliderInterval);
    };
  }, []);
  return (
    <div>
     <header className="header">
        <div className="row">
          <div className="col-md-12 col-sm-12 col-lg-12 header-sec">
            <div className="slider">
              {images.map((image, index) => (
                <div className={`slide ${index === currentIndex ? 'active' : ''}`} key={index}>
                  <img src={image} alt={`Slide ${index + 1}`} />
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
                <span className="text-btn-linkk">Book Now</span>
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
    </div>
  );
};

export default Home;

