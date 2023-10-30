import { useEffect,useState } from "react"
import React  from 'react'



    const Heading = () => {
        const [contant, setcontant] = useState({});
    
        const fetchDataFromAPI = () => {
            fetch('http://localhost:8080/admin/content/active')
                .then((res) => res.json())
                .then((data) => {
                    setcontant(data.data);
                });
        }
    
        useEffect(() => {
            fetchDataFromAPI();
        }, []);
        
        

  return (
   
        <div className="text-box ">
               <p className="text-p1">{contant.Heading}</p>
               <h3 className="text-h3">{contant.SubHeading}</h3>
                    <button className="text-button"><a href="#" className="text-btn-linkk">{contant.Button}</a></button>
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
   
  )
}

export default Heading;