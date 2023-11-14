import React, { useEffect, useState, useRef } from 'react';
import { Swiper, SwiperSlide } from 'swiper/react';
import 'swiper/swiper-bundle.css';
import 'swiper/css';
import 'swiper/css/pagination';
import './section2.css';

// import './styles.css';

// import required modules
import { Pagination } from 'swiper/modules';

function Section2() {
  const [testimonials, setTestimonials] = useState([]);

  useEffect(() => {
    const fetchTestimonials = async () => {
      try {
        const response = await fetch(
          'http://localhost:8080/admin/get/testimonials'
        );
        const data = await response.json();

        if (data.status === 200) {
          setTestimonials(data.data);
        } else {
          console.error('Error fetching testimonials:', data.message);
        }
      } catch (error) {
        console.error('Error fetching testimonials:', error);
      }
    };

    fetchTestimonials();

    // Initialize the Swiper after data is fetched
    //   const mySwiper = new Swiper('.mySwiper', {
    //     slidesPerView: 'auto',
    //     spaceBetween: 10,
    //     freeMode: true,
    //     navigation: {
    //       nextEl: '.swiper-button-next',
    //       prevEl: '.swiper-button-prev',
    //     },
    //     preventInteractionOnTransition: true,
    //     allowTouchMove: false,
    //     breakpoints: {
    //       767: {
    //         slidesPerView: 3,
    //         spaceBetween: 10,
    //       },
    //       576: {
    //         slidesPerView: 1,
    //         spaceBetween: 10,
    //       },
    //       // Add more breakpoints if needed
    //     },
    //   });

    //   const handleNextClick = () => {
    //     mySwiper.slideNext();
    //   };

    //   // Event handler for the Previous button
    //   const handlePrevClick = () => {
    //     mySwiper.slidePrev();
    //   };

    //   // Event listener for the Next and Previous buttons
    //   document
    //     .getElementById('swiper-button-next')
    //     .addEventListener('click', handleNextClick);

    //   document
    //     .getElementById('swiper-button-prev')
    //     .addEventListener('click', handlePrevClick);
  }, []); // Add an empty dependency array to ensure this effect runs only once

  return (
    <>
      <section className="bg-light py-5 py-xl-8 container">
        <div className="container">
          <div className="row">
            <div className="col-12 col-md-10 col-lg-8 col-xl-7 col-xxl-6">
              <h1 className="text-dark text-uppercase">
                Customers <b className="text-secondary">Testimonial</b>
              </h1>
              <hr className="w-100 mx-auto mb-5 mb-xl-9 border-dark-subtle" />
            </div>
          </div>
        </div>
        <Swiper
          slidesPerView={1}
          spaceBetween={30}
          pagination={{
            // clickable: true,
            dynamicBullets: true,
          }}
          breakpoints={{
            576: {
              slidesPerView: 1,
              spaceBetween: 10,
            },
            768: {
              slidesPerView: 2,
              spaceBetween: 20,
            },
            992: {
              slidesPerView: 3,
              spaceBetween: 30,
            },
          }}
          modules={[Pagination]}
        >
          <div className="container d-flex justify-content-center swiper-container swipeTestimonial">
            {testimonials.map((data) => (
              <SwiperSlide
                className="card border-0 border-bottom border-primary shadow-sm mb-3 me-4"
                key={data.ID}
              >
                <div className="card-body">
                  <figure>
                    <img
                      className="img-fluid rounded rounded-circle mb-4 border border-5 w-25"
                      loading="lazy"
                      src={`http://localhost:8080/admin/get/testimonial/image/${data.ID}`}
                      alt=""
                    />
                    <div
                      className="bsb-ratings text-warning mb-3"
                      data-bsb-star="5"
                      data-bsb-star-off="0"
                    ></div>
                    <blockquote className="bsb-blockquote-icon mb-4">
                      {data.Review}
                    </blockquote>
                    <h4 className="mb-2">{data.Name}</h4>
                    <h5 className="fs-6 text-secondary mb-0">
                      {data.Designation}
                    </h5>
                  </figure>
                </div>
              </SwiperSlide>
            ))}
            <div className="swiper-button-prev" id="swiper-button-prev"></div>
            <div className="swiper-button-next" id="swiper-button-next"></div>
          </div>{' '}
        </Swiper>
      </section>
    </>
  );
}

export default Section2;
