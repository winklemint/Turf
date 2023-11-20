import React, { useState } from 'react';
import { FaLocationDot } from 'react-icons/fa6';
import { FaStar } from 'react-icons/fa';
import { Swiper, SwiperSlide } from 'swiper/react';
import BookNowRules from './BookNowRules.js';
import Footer from './Footer.js';

import 'swiper/css';
import 'swiper/css/free-mode';
import 'swiper/css/navigation';
import 'swiper/css/thumbs';

import { FreeMode, Navigation, Thumbs } from 'swiper/modules';

function BookNow() {
  const h2Style = {
    color: 'rgb(119 196 40)',
  };

  const [thumbsSwiper, setThumbsSwiper] = useState(null);
  return (
    <>
      <div className="container">
        <div className="row float-center">
          <div className="col-sm-12 col-md-8 col-lg-8 col-12">
            <h1 className=" fw-1" style={h2Style}>
              Play The Turf, Rajendra Nagar - by Walking Dreamz
            </h1>
            <Swiper
              style={{
                '--swiper-navigation-color': '#fff',
                '--swiper-pagination-color': '#fff',
              }}
              loop={true}
              spaceBetween={10}
              navigation={true}
              thumbs={{ swiper: thumbsSwiper }}
              modules={[FreeMode, Navigation, Thumbs]}
              className="mySwiper2"
            >
              <SwiperSlide>
                <img
                  src="https://swiperjs.com/demos/images/nature-1.jpg"
                  style={{ width: '100%', height: '100%' }}
                />
              </SwiperSlide>
              <SwiperSlide>
                <img
                  src="https://swiperjs.com/demos/images/nature-2.jpg"
                  style={{ width: '100%', height: '100%' }}
                />
              </SwiperSlide>
              <SwiperSlide>
                <img
                  src="https://swiperjs.com/demos/images/nature-3.jpg"
                  style={{ width: '100%', height: '100%' }}
                />
              </SwiperSlide>
              <SwiperSlide>
                <img
                  src="https://swiperjs.com/demos/images/nature-4.jpg"
                  style={{ width: '100%', height: '100%' }}
                />
              </SwiperSlide>
              <SwiperSlide>
                <img
                  src="https://swiperjs.com/demos/images/nature-5.jpg"
                  style={{ width: '100%', height: '100%' }}
                />
              </SwiperSlide>
              <SwiperSlide>
                <img
                  src="https://swiperjs.com/demos/images/nature-6.jpg"
                  style={{ width: '100%', height: '100%' }}
                />
              </SwiperSlide>
              <SwiperSlide>
                <img
                  src="https://swiperjs.com/demos/images/nature-7.jpg"
                  style={{ width: '100%', height: '100%' }}
                />
              </SwiperSlide>
              <SwiperSlide>
                <img
                  src="https://swiperjs.com/demos/images/nature-8.jpg"
                  style={{ width: '100%', height: '100%' }}
                />
              </SwiperSlide>
              <SwiperSlide>
                <img
                  src="https://swiperjs.com/demos/images/nature-9.jpg"
                  style={{ width: '100%', height: '100%' }}
                />
              </SwiperSlide>
              <SwiperSlide>
                <img
                  src="https://swiperjs.com/demos/images/nature-10.jpg"
                  style={{ width: '100%', height: '100%' }}
                />
              </SwiperSlide>
            </Swiper>

            {/* <img src="/assets/football-ground-flooring.jpg" className="w-100" /> */}
          </div>
          <div className="col-sm-12 col-md-4 col-lg-4 col-12 mt-5 text-center ">
            <div className="fs-6 p-2 float-start">
              <div className="container" style={{ marginTop: '100px' }}>
                <div className="row">
                  <div className="col-1">
                    <FaLocationDot className="text-danger fs-3 me-3" />
                  </div>
                  <div className="col-11 text-center">
                    <p>
                      Rajendra Nagar, Near Rajendra Nagar Police station
                      Bijalpur, Indore (M.P.)
                    </p>
                  </div>
                  <div className="row text-start">
                    <div className="d-flex btn bg-warning text-light text-center col-2 badge text-wrap">
                      <FaStar className="me-2" /> 5
                    </div>
                    <div className="col-10 mb-1">
                      <div className="badge text-warning">
                        7 reviews / Write a review
                      </div>
                    </div>
                  </div>
                </div>
                <br />
                <div className="fs-6 text-start">Amenities:</div>
                <div className="fs-6 text-start">
                  <p
                    className="btn btn-secondary text-light me-2"
                    style={{ fontSize: '13px' }}
                  >
                    Seating
                  </p>
                  <p
                    className="btn btn-secondary text-light me-2"
                    style={{ fontSize: '13px' }}
                  >
                    Toilets
                  </p>
                  <p
                    className="btn btn-secondary text-light me-2"
                    style={{ fontSize: '13px' }}
                  >
                    Parking
                  </p>
                </div>
                <div className="text-start">
                  <button className="btn btn-outline-warning">Book Now</button>
                </div>
              </div>
            </div>
          </div>
        </div>
        {/* <BookNowRules /> */}
      </div>
      <BookNowRules />
      <Footer />
    </>
  );
}

export default BookNow;
