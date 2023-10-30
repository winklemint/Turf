
import React, { useEffect } from 'react';
import Carousel from 'react-bootstrap/Carousel';
import Swiper from 'swiper';
import './section1.css';

let Data = [
  {
    Id: 1,
    Palace: "Rajendra Nagar",
    Ground: "90 X 90",
    Amenities: "food, bev ",
    Img: "assets/59.jpg",
  },
  {
    Id: 2,
    Palace: "Vijay Nagar",
    Ground: "90 X 90",
    Amenities: "food, bev ",
    Img: "assets/turf-img.jpg",
  },
  {
    Id: 3,
    Palace: "Gandhi Nagar",
    Ground: "90 X 90",
    Amenities: "food, bev ",
    Img: "assets/football-ground-flooring.jpg",
  },
  {
    Id: 4,
    Palace: "IT Park",
    Ground: "90 X 90",
    Amenities: "food, bev ",
    Img: "assets/football-ground-flooring.jpg",
  },
];

function Section1() {
  useEffect(() => {
    // Inside your React component's useEffect hook
const mySwiper = new Swiper(".mySwiper", {
  slidesPerView: 'auto',
  spaceBetween: 10,
  centeredSlides: true,
  freeMode: true,
  loop: true,
  navigation: {
    nextEl: ".swiper-button-next",
    prevEl: ".swiper-button-prev",
  },
  noSwiping:true,
  noSwipingClass: "swiper-no-swiping",
  breakpoints: {
    767: {
      slidesPerView: 3,
      spaceBetween: 10,
    },
    576: {
      slidesPerView: 1,
      spaceBetween: 10,
    },
    // Add more breakpoints if needed
  },
});

mySwiper.allowTouchMove = false;

// Event listener for the Next button
document.getElementById("swiper-button-next").addEventListener("click", function () {
  mySwiper.slideNext();
});

// Event listener for the Previous button
document.getElementById("swiper-button-prev").addEventListener("click", function () {
  mySwiper.slidePrev();
});

  }, []); // Run this effect only once

  return (
    <section className="container slider-sec2">
      <div className="row">
        <div className="col-md-12 col-sm-12 col-lg-12">
          <div className="slider-sec2-heading">
            <p className="ex-p">EXCLUSIVELY</p>
            <p className="works-p">
              <span style={{ color: "purple", fontWeight: "bold" }}>works</span> with
            </p>
            <p className="start-p">Startups and founders</p>
          </div>
        </div>
        <div className="col-md-12 col-sm-12 col-lg-12">
          <div className="mySwiper">
            <div className="swiper-wrapper">
              {Data.map((data) => (
                <div className="swiper-slide" key={data.Id}>
                  <Carousel slide={false}>
                    <Carousel.Item>
                      <div className="content-med">
                        <div className="swiper-avatar">
                          <img src={data.Img} alt={data.Palace} />
                        </div>
                        <div className="cites-box">
                          <h2 className="cite">{data.Palace}</h2>
                          <p className="cite-box-parag">Ground Size - {data.Ground}</p>
                          <p className="cite-box-parag">Amenities - {data.Amenities}</p>
                          <button className="cite1">
                            <a href="#" className="btn-link">
                              Book Now →
                            </a>
                          </button>
                        </div>
                      </div>
                    </Carousel.Item>
                  </Carousel>
                </div>
              ))}
            </div>
            <div className="swiper-button-prev" id="swiper-button-prev"></div>
            <div className="swiper-button-next" id="swiper-button-next"></div>
          </div>
        </div>
      </div>
    </section>
  );
}

export default Section1;

