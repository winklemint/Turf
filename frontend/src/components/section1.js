import React, { useState, useEffect } from 'react';
import Swiper from 'swiper';
import 'swiper/swiper-bundle.css';
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
      freeMode: true,
      navigation: {
        nextEl: '.swiper-button-next',
        prevEl: '.swiper-button-prev',
      },
      preventInteractionOnTransition: true,
      allowTouchMove: false,
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

    const handleNextClick = () => {
      mySwiper.slideNext();
    };

    // Event handler for the Previous button
    const handlePrevClick = () => {
      mySwiper.slidePrev();
    };

    // Event listener for the Next and Previous buttons
    document
      .getElementById('swiper-button-next')
      .addEventListener('click', handleNextClick);

    document
      .getElementById('swiper-button-prev')
      .addEventListener('click', handlePrevClick);
  }, []); // Run this effect only once

  return (
    <section className="container slider-sec2">
      <div className="row">
        <div className="col-md-12 col-sm-12 col-lg-12">
          <div className="slider-sec2-heading ml-5">
            <p className="ex-p text-secondary fs-4">EXCLUSIVELY</p>
            <p className="works-p fs-1 fw-bold">
              <span className="font-weight-bold" style={{ color: 'purple' }}>
                works
              </span>{' '}
              with
            </p>
            <p className="start-p fs-1 fw-bold">Startups and founders</p>
          </div>
        </div>
        <div className="col-md-12 col-sm-12 col-lg-12">
          <div className="mySwiper swiper-container d-flex overflow-hidden">
            <div className="swiper-wrapper">
              {branchData.map((data) => (
                <div
                  className="swiper-slide text-center fs-5 align-items-center rounded border-0"
                  key={data.ID}
                >
                  <div className="content-med shadow-lg p-3 mb-5 bg-white rounded mt-1 me-5 col-sm-12 col-md-4 col-lg-12 ">
                    <div className="swiper-avatar">
                      <img
                        src={`http://localhost:8080/admin/branch/image/active/${data.ID}`}
                        alt={data.Turf_name}
                        className="img-fluid"
                      />
                    </div>
                    <div className="cites-box m-5">
                      <h2 className="cite fs-3 fw-bloder"> {data.Turf_name}</h2>
                      <p className="cite-box-parag">{data.Branch_name}</p>
                      <p className="cite-box-parag">
                        <i
                          className="fas fa-map-marker-alt"
                          style={{ color: 'red' }}
                        >
                          <span
                            className="address"
                            style={{ color: 'black', paddingLeft: '10px' }}
                          >
                            {data.Branch_address}
                          </span>
                        </i>
                      </p>
                      <button className="cite1 btn text-light rounded-pill border-0 fs-6 fw-bold p-2">
                        Book Now
                      </button>
                    </div>
                    <div className="sports-icon col-md-12 col-sm-12 col-lg-12 d-flex justify-content-center">
                      <span className="material-symbols-outlined tennis">
                        <img
                          className="sports-img"
                          src="assets/batminton.png"
                          alt="Badminton"
                        />
                      </span>
                      <span className="material-symbols-outlined cricket">
                        <img
                          className="sports-img"
                          src="assets/447875.png"
                          alt="Cricket"
                        />
                      </span>
                      <span className="material-symbols-outlined basketball">
                        <img
                          className="sports-img"
                          src="assets/footballllll.jpeg"
                          alt="Basketball"
                        />
                      </span>
                      <span className="material-symbols-outlined soccer">
                        <img
                          className="sports-img"
                          src="assets/fotbal123.png"
                          alt="Soccer"
                        />
                      </span>
                      <span className="material-symbols-outlined soccer">
                        <img
                          className="sports-img"
                          src="assets/tabletennis.png"
                          alt="Table Tennis"
                        />
                      </span>
                    </div>
                  </div>
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
