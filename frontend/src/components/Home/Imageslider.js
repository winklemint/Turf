import React, { useState, useEffect } from 'react';
import './ImageSlider.css'; // Import your CSS file

const ImageSlider = () => {
  const [currentIndex, setCurrentIndex] = useState(0);
  const [carouselImages, setCarouselImages] = useState([]);

  // Function to fetch carousel images
  const fetchCarouselImages = async () => {
    try {
      const response = await fetch('http://localhost:8080/admin/carousel/active');
      if (response.status === 200) {
        const responseData = await response.json();
        setCarouselImages(responseData.data);
      } else {
        throw new Error('Network response was not OK for carousel images');
      }
    } catch (error) {
      console.error('Error fetching carousel images: ' + error.message);
    }
  };

  // Function to advance to the next slide
  const nextSlide = () => {
    setCurrentIndex((currentIndex + 1) % carouselImages.length);
  };

  // UseEffect to fetch images and start the slider
  useEffect(() => {
    fetchCarouselImages();
  }, []);

  // UseEffect for automatic slider
  useEffect(() => {
    const sliderInterval = setInterval(() => {
      nextSlide();
    }, 3000);

    return () => {
      clearInterval(sliderInterval);
    };
  }, [currentIndex, carouselImages]);

  return (
    <div className="slider-container">
      
      <div className="slider">
        {carouselImages.map((image, index) => (
          <div className={`slide ${index === currentIndex ? 'active' : ''}`} key={index}>
            <img
              src={`http://localhost:8080/admin/get/image/active/${image.ID}`}
              alt={`Slide ${index + 1}`}
            />
          </div>
        ))}
      </div>
    </div>
  );
};

export default ImageSlider;
