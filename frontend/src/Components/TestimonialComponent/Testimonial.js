import React, { useEffect } from 'react';
import {  SwiperSlide } from 'swiper/react';
import './Testimonial.css';
import { Button, Card,  Col, Container,Row} from 'react-bootstrap';
import { Autoplay, Pagination } from 'swiper/modules';
import Swiper from 'swiper';

const Testimonial = () => {

    useEffect(() => {
      const testimonialSwiper = new Swiper('.swiper', {
        slidesPerView: 'auto', // Show as many slides as fit the container width
        spaceBetween: 30, // Space between slides
        centeredSlides: true, // Center the active slide
        autoplay: {
          delay: 3000, // Set the delay between slide transitions (in milliseconds)
        },
        loop: true,
        pagination: {
          el: '.swiper-pagination',
          clickable: true,
        },
        breakpoints: {
          767: {
            slidesPerView: 3,
            spaceBetween: 10,
          },
        },
      });
    }, []);
  return (
    <div>
      <section className="section3">
        <div className="container heading">
          <p className="client"><span style={{ color: '#ef1b40', fontWeight: 600 }}>Clients</span> memo</p>
          <p>testimonials</p>
          <p className="border-line"></p>
        </div>

        <div className="container container-swiper swiper">
          <div className="swiper-wrapper">
            <div className="swiper-slide">
              <div className="slide-box1">
                <h1 className="slide-box-h1">Perfect</h1>
                <p className="slide-box-p1">Lorem ipsum dolor sit amet, consectetur adipisicing elit, Ut enim ad</p>
                <p className="namp">- Rocky Hych</p>
                <h4 className="slide-box-h4">Founder, Arrow</h4>
              </div>
            </div>
            <div className="swiper-slide">
              <div className="slide-box2">
                <h1 className="slide-box-h1">Perfect</h1>
                <p className="slide-box-p1">Lorem ipsum dolor sit amet, consectetur adipisicing elit, Ut enim ad</p>
                <p className="namp">- Rocky Hych</p>
                <h4 className="slide-box-h4">Founder, Arrow</h4>
              </div>
            </div>
            <div className="swiper-slide">
              <div className="slide-box3">
                <h1 className="slide-box-h1">Perfect</h1>
                <p className="slide-box-p1">Lorem ipsum dolor sit amet, consectetur adipisicing elit, Ut enim ad</p>
                <p className="namp">- Rocky Hych</p>
                <h4 className="slide-box-h4">Founder, Arrow</h4>
              </div>
            </div>
            <div className="swiper-slide">
              <div className="slide-box4">
                <h1 className="slide-box-h1">Perfect</h1>
                <p className="slide-box-p1">Lorem ipsum dolor sit amet, consectetur adipisicing elit, Ut enim ad</p>
                <p className="namp">- Rocky Hych</p>
                <h4 className="slide-box-h4">Founder, Arrow</h4>
              </div>
            </div>
            <div className="swiper-slide">
              <div className="slide-box5">
                <h1 className="slide-box-h1">Perfect</h1>
                <p className="slide-box-p1">Lorem ipsum dolor sit amet, consectetur adipisicing elit, Ut enim ad</p>
                <p className="namp">- Rocky Hych</p>
                <h4 className="slide-box-h4">Founder, Arrow</h4>
              </div>
            </div>
          </div>
          <div className="swiper-pagination"></div>
        </div>
      </section>
    </div>
  );
};

export default Testimonial;
