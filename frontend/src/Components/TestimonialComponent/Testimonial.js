import React, { useEffect } from 'react';
import { Swiper, SwiperSlide } from 'swiper/react';
import './Testimonial.css';
import { Button, Card,  Col, Container,Row} from 'react-bootstrap';
import { Autoplay, Pagination } from 'swiper/modules';

const Testimonial = () => {

  return (
    <div>
      <section className="section3">
        <div className="container heading">
          <p className="client">
            <span style={{ color: "#ef1b40", fontWeight: 600 }}>Clients</span> memo
          </p>
          <p>testimonials</p>
          <p className="border-line"></p>
        </div>
        <Container className='swiper-container' >
    <Swiper
  slidesPerView={3}
  spaceBetween={40}
  pagination={{
    clickable: true,
  }}
  modules={[Autoplay, Pagination]} // Add Autoplay and Pagination modules
  autoplay={{ delay: 2000 }} // Set the delay (in milliseconds) between slides
  className="mySwiper"
>
           <SwiperSlide className='swip-slide' >
          <div className="testimonial-content">
          <div >
                    <h1 className='testimonial1'> Perfect </h1>
                    <p className='testimonial-p'>Lorem ipsum dolor sit amet, consectetur adipisicing elit, Ut enim ad </p>
                    <div class="text-right">
                      <p class="name">- John Doe</p>
                      <p class="designation">Founder, Arrow</p>
                    </div>
                  </div>
            </div></SwiperSlide>
        <SwiperSlide className='swip-slide1'>   <div className="testimonial-content">
          <div >
                    <h1 className='testimonial1'> Perfect </h1>
                    <p className='testimonial-p'>Lorem ipsum dolor sit amet, consectetur adipisicing elit,  Ut enim ad </p>
                    <div class="text-right">
                      <p class="name">- John Doe</p>
                      <p class="designation">Founder, Arrow</p>
                    </div>
                  </div>
            </div></SwiperSlide>
        <SwiperSlide className='swip-slide2' > 
        <div className="testimonial-content">
          <div >
                    <h1 className='testimonial1'> Perfect </h1>
                    <p className='testimonial-p'>Lorem ipsum dolor sit amet, consectetur adipisicing elit,  Ut enim ad </p>
                    <div class="text-right">
                      <p class="name">- John Doe</p>
                      <p class="designation">Founder, Arrow</p>
                    </div>
                  </div>
            </div>
            </SwiperSlide>
        <SwiperSlide className='swip-slide3'> 
        <div className="testimonial-content">
          <div >
                    <h1 className='testimonial1'>Perfect </h1>
                    <p className='testimonial-p'>Lorem ipsum dolor sit amet, consectetur adipisicing elit, Ut enim ad </p>
                    <div class="text-right">
                      <p class="name">- John Doe</p>
                      <p class="designation">Founder, Arrow</p>
                    </div>
                  </div>
            </div>
            </SwiperSlide>
        <SwiperSlide className='swip-slide4' > 
        <div className="testimonial-content">
          <div >
                    <h1 className='testimonial1'> Perfect </h1>
                    <p className='testimonial-p'>Lorem ipsum dolor sit amet, consectetur adipisicing elit,  Ut enim ad </p>
                    <div class="text-right">
                      <p class="name">- John Doe</p>
                      <p class="designation">Founder, Arrow</p>
                    </div>
                  </div>
            </div>
            </SwiperSlide>
        <SwiperSlide className='swip-slide5' > 
        <div className="testimonial-content">
          <div >
                    <h1 className='testimonial1'> Perfect </h1>
                    <p className='testimonial-p'>Lorem ipsum dolor sit amet,  Ut enim ad </p>
                    <div class="text-right">
                      <p class="name">- John Doe</p>
                      <p class="designation">Founder, Arrow</p>
                    </div>
                  </div>
            </div>
            </SwiperSlide>      </Swiper>
    </Container>
      </section>
    </div>
  );
};

export default Testimonial;
