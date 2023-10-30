import React, { useState, useEffect } from 'react';
import Carousel from 'react-bootstrap/Carousel';
import Swiper from 'swiper';
import './section1.css';

function Section1() {
  const [branchData, setBranchData] = useState([]);

  useEffect(() => {
    // Inside your React component's useEffect hook

    // Fetch data from the API
    const fetchBranchData = async () => {
      try {
        const response = await fetch(
          'http://localhost:8080/admin/active/branch'
        );
        if (!response.ok) {
          throw new Error('HTTP error: ' + response.status);
        }
        const data = await response.json();
        setBranchData(data.data);
      } catch (error) {
        console.error('Error fetching branch data:', error.message);
      }
    };

    fetchBranchData();

    // Initialize the Swiper after data is fetched
    const mySwiper = new Swiper('.mySwiper', {
      slidesPerView: 'auto',
      spaceBetween: 10,
      centeredSlides: true,
      freeMode: true,
      loop: true,
      navigation: {
        nextEl: '.swiper-button-next',
        prevEl: '.swiper-button-prev',
      },
      noSwiping: true,
      noSwipingClass: 'swiper-no-swiping',
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
    document
      .getElementById('swiper-button-next')
      .addEventListener('click', function () {
        mySwiper.slideNext();
      });

    // Event listener for the Previous button
    document
      .getElementById('swiper-button-prev')
      .addEventListener('click', function () {
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
              <span style={{ color: 'purple', fontWeight: 'bold' }}>works</span>{' '}
              with
            </p>
            <p className="start-p">Startups and founders</p>
          </div>
        </div>
        <div className="col-md-12 col-sm-12 col-lg-12">
          <div className="mySwiper">
            <div className="swiper-wrapper">
              {branchData.map((data) => (
                <div className="swiper-slide" key={data.ID}>
                  <Carousel slide={false}>
                    <Carousel.Item>
                      <div class="content-med ">
                        <div class="swiper-avatar">
                          <img
                            src={`http://localhost:8080/admin/branch/image/active/${data.ID}`}
                            alt={data.Turf_name}
                          />
                        </div>
                        <div class="cites-box">
                          <h2 class="cite"> {data.Turf_name}</h2>
                          <p className="cite-box-parag">{data.Branch_name}</p>
                          <p class="cite-box-parag">
                            <i
                              class="	fas fa-map-marker-alt"
                              style={{ color: 'red' }}
                            >
                              <span
                                class="address"
                                style={{ color: 'black', paddingLeft: '10px' }}
                              >
                                {data.Branch_address}
                              </span>
                            </i>
                          </p>
                          <button class="cite1">
                            <a href="#" class="btn-link">
                              Book Now{' '}
                            </a>
                          </button>
                        </div>
                        <div class="sports-icon">
                          <span class="material-symbols-outlined tennis">
                            <img
                              class="sports-img"
                              src="assets/batminton.png"
                            />
                          </span>
                          <span class="material-symbols-outlined cricket">
                            <img class="sports-img" src="assets/447875.png" />
                          </span>
                          <span class="material-symbols-outlined basketball">
                            <img
                              class="sports-img"
                              src="assets/footballllll.jpeg"
                            />
                          </span>
                          <span class="material-symbols-outlined soccer">
                            <img
                              class="sports-img"
                              src="assets/fotbal123.png"
                            />
                          </span>
                          <span class="material-symbols-outlined soccer">
                            <img
                              class="sports-img"
                              src="assets/tabletennis.png"
                            />
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
