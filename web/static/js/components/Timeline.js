import React from "react";
import PropTypes from "prop-types";
import { VerticalTimeline, VerticalTimelineElement } from "react-vertical-timeline-component";
import { BackHome } from "../BackHome/mod";
import { FaGithub } from "react-icons/fa";
import "react-vertical-timeline-component/style.min.css";

// import items from "./items";
// import img from "./bruneau.png";
// import london from "./london.jpeg";
// import lamoulade from "./lamoulade.png";
/**
 * @param {object} props
 * @param {{x:string, y:string}} props.home
 * @param {Function} props.setPosition
 */
export const Timeline = ({ setPosition, home }) => {
  //hacking the height at the moment
  // TODO fix styles
  return (
    <div style={{ overflowY: "scroll", height: "100vh" }}>
      <BackHome
        action={() => {
          const nextURl = "/#!home";
          const nextTitle = "Home | Abdoul Sy";
          const nextState = {
            additionalInformation: "moved from SmartParagraph to home page",
          };
          history.pushState(nextState, nextTitle, nextURl);
          document.title = nextTitle;
          setPosition(home);
        }}
      />
      <VerticalTimeline>
        <VerticalTimelineElement
          className="vertical-timeline-element--work"
          contentStyle={{ background: "rgb(33, 150, 243)", color: "#fff" }}
          contentArrowStyle={{ borderRight: "7px solid  rgb(33, 150, 243)" }}
          date="Jan 2020 - May 2020"
          iconStyle={{ background: "rgb(33, 150, 243)", color: "#fff" }}
          icon={<FaGithub />}
        >
          <h3 className="vertical-timeline-element-title">MrAndMrsSmith - Senior Engineer</h3>
          <h4 className="vertical-timeline-element-subtitle">London, UK</h4>
          <p>
            Lead Frontend for two scrum teams Created a dockerized stack for front-end development optimized for
            microservices. In the process of rewriting the website using React, Typescript, GraphQL, next js. nodeJS
            Introduced Contract Driven Development alongside Architectural Decision Records Created TDD and Integration
            testing workshops Introduced Performance testing for the new blog stack Introduced Behaviour Driven
            Development testing with Cypress. Integration with Jira, Hipchat, Gitlab, Github, and Test Report Frameworks
            (Testrail like) Performing management meeting (1-1’s team lead meetings) Symfony 2 development on the
            backend. Using Ontologies to link the domain of the company.
          </p>
        </VerticalTimelineElement>
        <VerticalTimelineElement
          className="vertical-timeline-element--work"
          contentStyle={{ background: "rgb(33, 150, 243)", color: "#fff" }}
          contentArrowStyle={{ borderRight: "7px solid  rgb(33, 150, 243)" }}
          date="2017 - 2019"
          iconStyle={{ background: "rgb(33, 150, 243)", color: "#fff" }}
          icon={<FaGithub />}
        >
          <h3 className="vertical-timeline-element-title">Gamesys - Senior Javascript Developer</h3>
          <h4 className="vertical-timeline-element-subtitle">London, UK</h4>
          <p>
            Lead a team of javascript developers for the new Full Stack application Created an Oauth2 login system
            wrapping OpenId and SAML applications Developed the local docker stack for development in https Worked on an
            embedded chat application using websockets XMPP technology stack Video Streaming stack with wowza Security
            Champion of the team Created a Twistlock integration framework Created a Sonarcube PR decorator Taught TDD,
            Docker, Gocd, in workshops Using terraform to spin up development environments
          </p>
        </VerticalTimelineElement>
        <VerticalTimelineElement
          className="vertical-timeline-element--work"
          contentStyle={{ background: "rgb(33, 150, 243)", color: "#fff" }}
          date="2014 - 2016"
          iconStyle={{ background: "rgb(33, 150, 243)", color: "#fff" }}
          icon={<FaGithub />}
        >
          <h3 className="vertical-timeline-element-title">Smartlogic - Lead Web Developer</h3>
          <h4 className="vertical-timeline-element-subtitle">London, UK</h4>
          <p>
            Lead a team of developers in Poland for the new cloud application - Semaphore. Developed the Smartlogic.com,
            on my own with the Lead UX designer, 12 times faster than the previous team of 10 (1 month). Author and
            Architect of a complex D3 Visualizer. Implemented the continuous development and testing practices for the
            Frontend part of the application. Created the architecture of the frontend part of the newest Semaphore
            Application. ● Attended various training sessions on RDF technologies, giving presentations.
          </p>
        </VerticalTimelineElement>
        <VerticalTimelineElement
          className="vertical-timeline-element--work"
          date="2013 - 2014"
          contentStyle={{ background: "rgb(33, 150, 243)", color: "#fff" }}
          iconStyle={{ background: "rgb(33, 150, 243)", color: "#fff" }}
          icon={<FaGithub />}
        >
          <h3 className="vertical-timeline-element-title">Byng Systems, Senior Interactions Developer</h3>
          <h4 className="vertical-timeline-element-subtitle">London, UK</h4>
          <p>
            Took over the development of the App Side of an iPad App: (Consentz). ( Backbone.js and Cordova Phonegap).
            Wrote various SVG heavy websites, and heavy animated websites. Wrote documentation and tried to implement
            good testing practices.
          </p>
        </VerticalTimelineElement>
        <VerticalTimelineElement
          className="vertical-timeline-element--work"
          date="2010 - 2013"
          contentStyle={{ background: "rgb(33, 150, 243)", color: "#fff" }}
          iconStyle={{ background: "rgb(33, 150, 243)", color: "#fff" }}
          icon={<FaGithub />}
        >
          <h3 className="vertical-timeline-element-title">La Moulade - Senior Web Developer</h3>
          <h4 className="vertical-timeline-element-subtitle">London, UK</h4>
          <p>
            Designed all the technical architectures for the web application for our clients. Wrote complex Js based
            applications. Developed PCI compliant websites and compatible for IE6+ without becoming crazy. Documented
            the technical side of the business plan for Juun. Won several awards for excellent JS and CSS on
            Lamoulade.com creating unique algorithms for what would become superscrollorama ( CSS Awards, and AWWWards
            site of the day - runner up for Site of the year).
          </p>
        </VerticalTimelineElement>
        <VerticalTimelineElement
          className="vertical-timeline-element--education"
          date="2008 - 2009"
          contentStyle={{ background: "rgb(33, 150, 243)", color: "#fff" }}
          iconStyle={{ background: "rgb(233, 30, 99)", color: "#fff" }}
          icon={<FaGithub />}
        >
          <h3 className="vertical-timeline-element-title">Webmaster</h3>
          <h4 className="vertical-timeline-element-subtitle">JM-Bruneau</h4>
          <p>
            Creating promotional HTML banners / Emails Created a CMS to automate creating Emails and HTML banners from
            the assets the retailers would give. Upgraded the CMS to automate asset organisations Ab-Tool the CMS was
            able to generate an HTML in less than a second when it would take a day before (~21 000 times faster) Could
            organize assets with a 0.5 margin of error in less than a 10 seconds, when it would take days to organise
            assets (big number faster)
          </p>
        </VerticalTimelineElement>
        <VerticalTimelineElement
          className="vertical-timeline-element--education"
          date="2009"
          contentStyle={{ background: "rgb(33, 150, 243)", color: "#fff" }}
          iconStyle={{ background: "rgb(233, 30, 99)", color: "#fff" }}
          icon={<FaGithub />}
        >
          <h3 className="vertical-timeline-element-title">Zend Certified PHP Developer</h3>
          <h4 className="vertical-timeline-element-subtitle">Certification</h4>
          <p>PHP 5.2 Development</p>
        </VerticalTimelineElement>
        <VerticalTimelineElement
          className="vertical-timeline-element--education"
          date="2009"
          contentStyle={{ background: "rgb(33, 150, 243)", color: "#fff" }}
          iconStyle={{ background: "rgb(233, 30, 99)", color: "#fff" }}
          icon={<FaGithub />}
        >
          <h3 className="vertical-timeline-element-title">MySQL Certified Associate</h3>
          <h4 className="vertical-timeline-element-subtitle">Certification</h4>
          <p>MySQL 5.x</p>
        </VerticalTimelineElement>
        <VerticalTimelineElement
          className="vertical-timeline-element--education"
          date="2008 - 2009"
          contentStyle={{ background: "rgb(33, 150, 243)", color: "#fff" }}
          iconStyle={{ background: "rgb(233, 30, 99)", color: "#fff" }}
          icon={<FaGithub />}
        >
          <h3 className="vertical-timeline-element-title">Bachelor of Information Technology in Web Development</h3>
          <h4 className="vertical-timeline-element-subtitle">Bachelor Degree</h4>
          <p>Web Development</p>
        </VerticalTimelineElement>
        <VerticalTimelineElement iconStyle={{ background: "rgb(16, 204, 82)", color: "#fff" }} icon={<FaGithub />} />
      </VerticalTimeline>
    </div>
  );
};

Timeline.propTypes = {
  home: PropTypes.shape({
    x: PropTypes.string.isRequired,
    y: PropTypes.string.isRequired,
  }).isRequired,
  setPosition: PropTypes.func.isRequired,
};
