import React, { useState, useEffect } from 'react';
import 'swiper/swiper-bundle.css';
import SwiperCore, { Pagination } from 'swiper';
import { Swiper, SwiperSlide } from 'swiper/react';

SwiperCore.use([Pagination]);

const ImageSlider = () => {
  const [swiper, setSwiper] = useState(null);
  const [carouselData, setCarouselData] = useState([]);

  const handleSlideChange = (swiper) => {
    // Handle slide change
  };

  const goToSlide = (index) => {
    if (swiper) {
      swiper.slideTo(index);
    }
  };

  // Fetch carousel data from your API
  useEffect(() => {
    // Replace with your API endpoint
    fetch('your_api_endpoint_here')
      .then((response) => response.json())
      .then((data) => setCarouselData(data))
      .catch((error) => console.error('Error fetching carousel data:', error));
  }, []);

  return (
    <div>
      <Swiper
        onSwiper={setSwiper}
        onSlideChange={handleSlideChange}
        pagination={{ clickable: true }}
      >
        {carouselData.map((item, index) => (
          <SwiperSlide key={index}>
            {/* Render slide content with data from the API */}
            <img src={item.image} alt={`Slide ${index + 1}`} />
          </SwiperSlide>
        ))}
      </Swiper>
      <div className="custom-pagination">
        {carouselData.map((_, index) => (
          <span key={index} onClick={() => goToSlide(index)} className="pagination-dot"></span>
        ))}
      </div>
    </div>
  );
};

export default ImageSlider;
