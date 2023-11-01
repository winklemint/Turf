
import React,{useState,useEffect} from "react";
import { Swiper, SwiperSlide } from 'swiper/react';
import 'swiper/swiper-bundle.css';

function Section2() {
    const [carouselData, setCarouselData] = useState([]);

    useEffect(() => {
        fetch('http://localhost:8080/admin/get/testimonials')
          .then((response) => response.json())
          .then((data) => setCarouselData(data.data))
          .catch((error) => console.error('Error fetching carousel data:', error));
      }, []);

    return (
        <section className="section3">
            <div className="container heading">
                <p className="client"><span style={{ color: "#ef1b40", fontWeight: "600" }}>Clients</span> memo</p>
                <p>testimonials</p>
                <p className="border-line"></p>
            </div>

            <div className="container container-swiper swiper">
                <div className="swiper-wrapper">
                    <Swiper
                        spaceBetween={50}
                        onSlideChange={() => console.log('slide change')}
                        onSwiper={(swiper) => console.log(swiper)}
                        mousewheel={{ forceToAxis: true }}
                        pagination={{ clickable: true }}
                        breakpoints={{
                            768: {
                                slidesPerView: 3, // Number of slides on screens wider than 768px
                            },
                            576: {
                                slidesPerView: 1, // Number of slides on screens wider than 576px
                            },
                            0: {
                                slidesPerView:.5, // Number of slides on screens at or smaller than 576px
                            },
                        }}
                    >
                                {carouselData.map((item,index) => (
          <       SwiperSlide key={index}>
                 {/* Render slide content with data from the API */}
                 <div className="swiper-slide">
                                <div className="slide-box1">
                                    <h1>{item.Designation}</h1>
                                    <p>{item.Review}</p>
                                    <p className="namp">{item.Name}</p>
                                </div>
                            </div>
                 </SwiperSlide>
                                ))}
                        
                    </Swiper>
                </div>
            </div>
        </section>
    )
}

export default Section2;


