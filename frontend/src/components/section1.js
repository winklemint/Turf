
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
                                  <div class="content-med ">
                                    <div class="swiper-avatar"><img src={data.Img}/></div>
                                      <div class="cites-box">
                                    <h2 class="cite"> {data.Palace}</h2>
                                    <p class="cite-box-parag"><i class="	fas fa-map-marker-alt" style= {{color: "red"}} ><span class="address" style={{color:"black", paddingLeft:"10px"}}>Indore (M.P)</span></i></p>
                                    <button class="cite1"><a href="#" class="btn-link">Book Now </a></button>
                                </div>
                                <div class="sports-icon">
                              
                                
                                <span class="material-symbols-outlined tennis">
                                <img class="sports-img"
                                src="assets/batminton.png"/>
                                </span>
                                <span class="material-symbols-outlined cricket">
                                <img class="sports-img"
                                src="assets/447875.png"/>
                                </span>
                                <span class="material-symbols-outlined basketball">
                                <img class="sports-img"
                                src="assets/footballllll.jpeg"/>
                                </span>
                                <span class="material-symbols-outlined soccer">
                                <img class="sports-img"
                                src="assets/fotbal123.png"/>
                                </span>
                                <span class="material-symbols-outlined soccer">
                                <img class="sports-img"
                                src="assets/tabletennis.png"/>
                                </span>
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

