import React from "react";
import { Swiper, SwiperSlide } from 'swiper/react';
import 'swiper/css';



function Section2(){
    return(
        <section className="section3">
			<div className="container heading">
				<p className="client"><span style={{color:"#ef1b40",fontWeight: "600"}}>Clients</span> memo</p>
				<p>testimonials</p>
				<p className="border-line"></p>
				
			</div>

			<div className=" container container-swiper swiper">
	            <div className=" swiper-wrapper">
				<Swiper
					spaceBetween={50}
					slidesPerView={3}
					onSlideChange={() => console.log('slide change')}
					onSwiper={(swiper) => console.log(swiper)}
					>
					<SwiperSlide>
						<div className="swiper-slide">
	            		<div className="slide-box1">
	            		<h1>Built with Bootstrap</h1>
	            		<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor </p>
	            		<p className="namp">Rocky Hych</p>
	            	    </div>
	            	</div>
					</SwiperSlide>
					<SwiperSlide>
						<div className="swiper-slide">
	            		<div className="slide-box2">
		            		<h1>Built with Bootstrap</h1>
		            		<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor </p>
		            		<p className="namp">Rocky Hych</p>
		            	</div>	
	            	</div>
					</SwiperSlide>
					<SwiperSlide>
						<div className="swiper-slide">
	            		<div className="slide-box3">
		            		<h1>Built with Bootstrap</h1>
		            		<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor </p>
		            		<p className="namp">Rocky Hych</p>
		            	</div>	 
	            	</div>
					</SwiperSlide>
					<SwiperSlide>
						<div className="swiper-slide">
	            		<div className="slide-box4">
		            		<h1>Built with Bootstrap</h1>
		            		<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor </p> 
		            		<p className="namp">Rocky Hych</p>
		            	</div>	
	            	</div>
					</SwiperSlide>
					<SwiperSlide>
						<div className="swiper-slide">
	            		<div className="slide-box5">
		            		<h1>Built with Bootstrap</h1>
		            		<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor </p>
		            		<p className="namp">Rocky Hych</p>
		            	</div>	
	            	</div>
					</SwiperSlide>
					...
					</Swiper>
	            	
	            	
	            	
	            	
	            	
	            
	            	       {/* <div className="swiper-pagination"></div>  */}
	        
	            </div>
                   
     
            </div>
			
		</section>
    )
}

export default Section2;