import React from 'react';

function AquaticAnimals() {
  const animalData = [
    {
      title: 'Jellyfish',
      description:
        'Jellyfish and sea jellies are the informal common names given to the medusa-phase of certain gelatinous members of the subphylum Medusozoa, a major part of the phylum Cnidaria.',
      link: 'https://en.wikipedia.org/wiki/Jellyfish',
    },
    {
      title: 'Seahorse',
      description:
        'Seahorses are mainly found in shallow tropical and temperate salt water throughout the world. They live in sheltered areas such as seagrass beds, estuaries, coral reefs, and mangroves. Four species are found in Pacific waters from North America to South America.',
      link: 'https://en.wikipedia.org/wiki/Seahorse',
    },
    {
      title: 'Octopus',
      description:
        'Octopuses inhabit various regions of the ocean, including coral reefs, pelagic waters, and the seabed; some live in the intertidal zone and others at abyssal depths. Most species grow quickly, mature early, and are short-lived.',
      link: 'https://en.wikipedia.org/wiki/Octopus',
    },
    {
      title: 'Shark',
      description:
        'Sharks are a group of elasmobranch fish characterized by a cartilaginous skeleton, five to seven gill slits on the sides of the head, and pectoral fins that are not fused to the head.',
      link: 'https://en.wikipedia.org/wiki/Shark',
    },
    {
      title: 'Dolphin',
      description:
        'Dolphins are widespread. Most species prefer the warm waters of the tropic zones, but some, such as the right whale dolphin, prefer colder climates. Dolphins feed largely on fish and squid, but a few, such as the orca, feed on large mammals such as seals.',
      link: 'https://en.wikipedia.org/wiki/Dolphin',
    },
  ];

  return (
    <main>
      <div>
        <span>discover</span>
        <h1>aquatic animals</h1>
        <hr />
        <p>
          Beauty and mystery are hidden under the sea. Explore with our
          application to know about Aquatic Animals.
        </p>
        <a href="#">download app</a>
      </div>
      <div className="swiper">
        <div className="swiper-wrapper">
          {animalData.map((animal, index) => (
            <div
              key={index}
              className={`swiper-slide swiper-slide--${index + 1}`}
            >
              <div>
                <h2>{animal.title}</h2>
                <p>{animal.description}</p>
                <a href={animal.link} target="_blank" rel="noopener noreferrer">
                  explore
                </a>
              </div>
            </div>
          ))}
        </div>
      </div>
      <img
        src="https://cdn.pixabay.com/photo/2021/11/04/19/39/jellyfish-6769173_960_720.png"
        alt=""
        className="bg"
      />
      <img
        src="https://cdn.pixabay.com/photo/2012/04/13/13/57/scallop-32506_960_720.png"
        alt=""
        className="bg2"
      />
    </main>
  );
}

export default AquaticAnimals;
