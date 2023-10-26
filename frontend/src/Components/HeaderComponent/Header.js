import React, { useEffect } from 'react';
import './Header.css';
import {Link} from 'react-router-dom';

const Header = () => {
 
    useEffect(() => {
        const slider = document.querySelector('.slider');
        const slides = document.querySelectorAll('.slide');
        let currentIndex = 0;
    
        function nextSlide() {
          currentIndex = (currentIndex + 1) % slides.length;
          updateSlider();
        }
    
        function updateSlider() {
          slider.style.transform = `translateX(-${currentIndex * 100}%)`;
        }
    
        function startSlider() {
          setInterval(nextSlide, 3000);
        }
    
        startSlider();
      }, []);
  return (
    <div>
      <header className="header">
      <div class="row">	
		   <div class="col-md-12 col-sm-12 col-lg-12 header-sec">
            <div class="slider">
	            <div class="slide"><img src="assets/images/birthdayparties2.jpg"/></div>
	            <div class="slide"><img src=" assets/images/DSC00867.jpg"/></div>
	            <div class="slide"><img src="assets/images/sixer-zon.jpg"/></div>
            </div>
         </div>
			
	           
	     	<div  class="col-md-12 col-sm-4 col-lg-12">
	           		<nav>
	           <div class="menu-toggle" id="menuToggle">
	            <div class="hamburger">
	                <span class="line"></span>
	                <span class="line"></span>
	                <span class="line-last"></span>
	            </div>
	            <div class="cross" style={{display: "none"}}>
	                <span class="line-cross"></span>
	                <span class="line-cross1"></span>
	            </div>
	          </div>
	          <ul class="main-nav" id="mainNav" style={{display: "none"}}>
            <li>
  <Link to="/">Home</Link>
</li>
<li>
  <Link to="/booking">Booking</Link>
</li>
<li>
  <Link to="/about">About</Link>
</li>
<li>
  <Link to="/contact">Contact</Link>
</li>
	          </ul>
	          </nav>
	       </div>
	       </div>
           <div class="col-md-12 col-sm-12 col-lg-12">

<div class="text-box ">
    <p class="text-p1">Hii ,I am</p>
    <h3 class="text-h3">Turf</h3>
    <p class="text-p2">Freelance designer from Melbourne</p>
    <button class="text-button"><a href="#" class="text-btn-linkk">Book Now</a></button>
    <div class="icon-sec">
        <p>Join me here</p>
        <div class="brand-icon">
        <i class='fab fa-whatsapp' style={{fontsize:"30px"}}></i>
        <i class='fab fa-facebook' style={{fontsize:"30px"}}></i>
        <i class='fab fa-twitter' style={{fontsize:"30px"}}></i>
        <i class='fab fa-linkedin' style={{fontsize:"30px"}}></i>
        </div>
</div>	   

    </div>
</div>

      </header>
   
    </div>
  );
};

export default Header;
