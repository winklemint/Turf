import React, { useEffect, useState } from 'react';
  
const Header = () => {
            const [currentIndex, setCurrentIndex] = useState(0);
            useEffect(() => {
                const slider = document.querySelector('.slider');
                const slides = document.querySelectorAll('.slide');
                const nextSlide = () => {
                    setCurrentIndex((currentIndex + 1) % slides.length);
                    slider.style.transform = `translateX(-${currentIndex * 100}%)`;
                };
                const sliderInterval = setInterval(nextSlide, 3000);
                return () => {
                    clearInterval(sliderInterval);
                };
            }, [currentIndex]);

        const menubar = () =>{
            const menu = document.getElementById('mainNav');
            const cross = document.getElementById('cross');
            const tongle = document.getElementsByClassName('hamburger');
            
            menu.style.display ='block';
            cross.style.display = 'block';
            if (tongle.length > 0) {
                for (let i = 0; i < tongle.length; i++) {
                    tongle[i].style.display = 'none';
                }}


        }

        const menubarhide = () =>{
            const menu = document.getElementById('mainNav');
            const cross = document.getElementById('cross');
            const tongle = document.getElementsByClassName('hamburger');
            
            menu.style.display ='none';
            cross.style.display = 'none';
            if (tongle.length > 0) {
                for (let i = 0; i < tongle.length; i++) {
                    tongle[i].style.display = 'block';
                }}

        }

       


    const [data, setData] = useState([]);

    useEffect(() => {
        fetchData();
    }, []);

  const fetchData = async () => {
    try {
      const response = await fetch('https://api.example.com/data'); // Replace with your API URL
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      const jsonData = await response.json();
      setData(jsonData);
    } catch (error) {
      console.error('Error fetching data:', error);
    }
  };




                
        

    return(
        <div className="App">

            <div className=" container-fuild">
            <header className=" header"/>
            <div className="row">	
                <div className="col-md-12 col-sm-12 col-lg-12 header-sec">
                    <div className="slider">
                        <div className="slide"><img src="assets/birthdayparties2.jpg"/></div>
                        <div className="slide"><img src="assets/DSC00867.jpg"/></div>
                        <div className="slide"><img src="assets/sixer-zon.jpg"/></div>
                    </div>
                    </div>
                
                        
                <div  className="col-md-12 col-sm-4 col-lg-12">
                        <nav>
                        <div className="menu-toggle" id="menuToggle">
                        <div className="hamburger" onClick={menubar} style={{display:"block"}}>
                            <span className="line"></span>
                            <span className="line"></span>
                            <span className="line-last"></span>
                        </div>
                        <div className="cross" id='cross' style={{display: "none"}} onClick={menubarhide}>
                            <span className="line-cross"></span>
                            <span className="line-cross1"></span>
                        </div>
                    </div>
                    <ul className="main-nav" id="mainNav" style={{display: "none",marginRight:"105px"}}>
                        <li><a href="#">Home</a></li>
                        <li><a href="#">Portfolio</a></li>
                        <li><a href="#">About</a></li>
                        <li><a href="#">Contact</a></li>
                    </ul>
                    </nav>
                    </div>
                    
                    




                    <div className="col-md-12 col-sm-12 col-lg-12">

                <div className="text-box ">
                    <p className="text-p1">Hii ,I am</p>
                    <h3 className="text-h3">Turf</h3>
                    <p className="text-p2">Freelance designer from Melbourne</p>
                    <button className="text-button"><a href="#" className="text-btn-linkk">Book Now</a></button>
                    <div className="icon-sec">
                    <p>Join me here</p>
                    <div className="brand-icon">
                    <i className='fab fa-whatsapp' style={{fontSize:"30px"}}></i>
                    <i className='fab fa-facebook' style={{fontSize:"30px"}}></i>
                    <i className='fab fa-twitter' style={{fontSize:"30px"}}></i>
                    <i className='fab fa-linkedin' style={{fontSize:"30px"}}></i>
                        </div>
                </div>	   

                    </div>
                </div>
             </div>
        </div>
    </div>
    );
        

    }


export default Header;