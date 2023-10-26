import React from 'react'

 const Swiper = () => {
  return (
    <div><section className="container slider-sec2">
    <div className="row">
      <div className="col-md-12 col-sm-12 col-lg-12">
        <div className="slider-sec2-heading">
          <p className="ex-p">EXCLUSIVELY</p>
          <p className="works-p"><span style={{ color: "purple", fontWeight: "bold" }}>works</span> with</p>
          <p className="start-p">Startups and founders</p>
          <p></p>
        </div>
      </div>
      <div className="col-md-12 col-sm-12 col-lg-12">
        <div className="mySwiper">
          <div className="swiper-wrapper">
            <div className="swiper-slide">
              <div className="content-med">
                <div className="swiper-avatar"><img src="assets/images/59.jpg" alt="Slide 1" /></div>
                <div className="cites-box">
                  <h2 className="cite">Rajenndra Nagar</h2>
                  <p className="cite-box-parag">Ground size: 90x90</p>
                  <p className="cite-box-parag">Amenities: food, bev</p>
                  <button className="cite1"><a href="#" className="btn-link">Book Now →</a></button>
                </div>
              </div>
            </div>
            <div className="swiper-slide">
              <div className="content-med">
                <div className="swiper-avatar"><img src="assets/images/turf-img.jpg" alt="Slide 2" /></div>
                <div className="cites-box">
                  <h2 className="cite">Vijay Nagar</h2>
                  <p className="cite-box-parag">Ground size: 90x90</p>
                  <p className="cite-box-parag">Amenities: food, bev</p>
                  <button className="cite1"><a href="#" className="btn-link">Book Now →</a></button>
                </div>
              </div>
            </div>
            <div className="swiper-slide">
              <div className="content-med">
                <div className="swiper-avatar"><img src="assets/images/football-ground-flooring.jpg" alt="Slide 3" /></div>
                <div className="cites-box">
                  <h2 className="cite">IT Park</h2>
                  <p className="cite-box-parag">Ground size: 90x90</p>
                  <p className="cite-box-parag">Amenities: food, bev</p>
                  <button className="cite1"><a href="#" className="btn-link">Book Now →</a></button>
                </div>
              </div>
            </div>
            <div className="swiper-slide">
              <div className="content-med">
                <div className="swiper-avatar"><img src="assets/images/turf-img.jpg" alt="Slide 4" /></div>
                <div className="cites-box">
                  <h2 className="cite">Gandhi Nagar</h2>
                  <p className="cite-box-parag">Ground size: 90x90</p>
                  <p className="cite-box-parag">Amenities: food, bev</p>
                  <button className="cite1"><a href="#" className="btn-link">Book Now →</a></button>
                </div>
              </div>
            </div>
            <div className="swiper-slide">
              <div className="content-med">
                <div className="swiper-avatar"><img src="assets/images/football-ground-flooring.jpg" alt="Slide 5" /></div>
                <div className="cites-box">
                  <h2 className="cite">Annapurna</h2>
                  <p className="cite-box-parag">Ground size: 90x90</p>
                  <p className="cite-box-parag">Amenities: food, bev</p>
                  <button className="cite1"><a href="#" className="btn-link">Book Now →</a></button>
                </div>
              </div>
            </div>
          </div>
          <div className="swiper-button-prev"></div>
          <div className="swiper-button-next"></div>
        </div>
      </div>
    </div>
  </section></div>
  )
}
export default Swiper;