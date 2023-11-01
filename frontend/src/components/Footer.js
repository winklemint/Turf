import React from "react";
 


function Footer(props){
    return(

        <footer className="footer">
			<div className="footer-box">
				<p className="footer-p1">{props.Headingdata.Footer}</p>
				<p className="footerp"></p>
						<div className="foot-brand-icon">
								<i className='fab fa-whatsapp' style={{fontSize:"45px"}}></i>
								<i className='fab fa-facebook' style={{fontSize:"45px"}}></i>
								<i className='fab fa-twitter' style={{fontSize:"45px"}}></i>
								<i className='fab fa-linkedin' style={{fontSize:"45px"}}></i>
					    </div>
			</div>
			<div className="foot-copyright">
             <p>Designed by <span style={{color:"#c2c2c2"}}>walkingdreamz</span> | Download <span style={{color: "#c2c2c2"}}>walkingdreamz</span></p>
		    </div>
		</footer>

    );
}

export default Footer;