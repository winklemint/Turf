import React from 'react'
import { useState, useEffect } from 'react';
import Swiper from 'swiper';

const SwiperComponent = () => {
	const [branches, setBranches] = useState([]);
  
	const fetchData = async () => {
	  try {
		const response = await fetch(`http://localhost:8080/admin/active/branch`);
		console.log('Response Status:', response.status);
		if (response.status === 201) {
		  const responseData = await response.json();
		  console.log("Branch API response", responseData);
		  setBranches(responseData.data); // Use the correct property for the array
		} else {
		  throw new Error('Network response was not ok');
		}
	  } catch (error) {
		console.error('Error fetching Branches: ' + error.message);
	  }
	};
  
	useEffect(() => {
	  fetchData();
	}, []);
	useEffect(() => {
		const testimonialSwiper = new Swiper('.swiper', {
		  slidesPerView: 'auto', // Show asmany slides as fit the container width
		  spaceBetween: 30, // Space between slides
		  centeredSlides: true, // Center the active slide
		  autoplay: {
			delay: 2000, // Set the delay between slide transitions (in milliseconds)
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
	  }, []); // Empty dependency array means this will only run once on component mount
	
  return (
	<div>
	<section className="container slider-sec2">
	  <div>
		<div>
		  <div className="slider-sec2-heading">
			<p className="ex-p">EXCLUSIVELY</p>
			<p className="works-p">
			  <span style={{ color: "purple", fontWeight: "bold" }}>works</span> with
			</p>
			<p className="start-p">Startups and founders</p>
			<p></p>
		  </div>
		</div>
	  </div>
	  <div>
	  <div className="mySwiper">
            <div className="swiper-wrapper">
              {branches.map((branch, index) => (
                <div className="swiper-slide" key={index}>
                  <div className="content-med">
                    <div className="swiper-avatar">
                    <img
                        src={`http://localhost:8080/admin/branch/image/active/${branch.ID}`}
                        alt={branch.Turf_name}
                      />
                    </div>
                    <div className="cites-box">
                      <h2 className="cite">{branch.Branch_name}</h2>
                      <p className="cite-box-parag">
                        <i className="fas fa-map-marker-alt" style={{ color: "red" }}>
                          <span className="address" style={{ color: "black", paddingLeft: "10px" }}>
                            {branch.Branch_address}
                          </span>
                        </i>
                      </p>
                      <button className="cite1">
                        <a href="#" className="btn-link">
                          Book Now
                        </a>
                      </button>
				  </div>
				  <div className="sports-icon">
					<span className="material-symbols-outlined tennis">
					  <img className="sports-img" src="assets/images/batminton.png" />
					</span>
					<span className="material-symbols-outlined cricket">
					  <img className="sports-img" src="assets/images/447875.png" />
					</span>
					<span className="material-symbols-outlined basketball">
					  <img className="sports-img" src="assets/images/footballllll.jpeg" />
					</span>
					<span className="material-symbols-outlined soccer">
					  <img className="sports-img" src="assets/images/fotbal123.png" />
					</span>
					<span className="material-symbols-outlined soccer">
					  <img className="sports-img" src="assets/images/tabletennis.png" />
					</span>
				  </div>
				</div>
			  </div>
			))}
		  </div>
		  <div className="swiper-button-prev"></div>
		  <div className="swiper-button-next"></div>
		</div>
	  </div>
	</section>
  </div>
  )
}
export default  SwiperComponent;