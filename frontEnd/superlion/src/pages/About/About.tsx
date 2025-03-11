import { useEffect, useState } from "react";

import ori_1 from "./ori1.jpg";
import ori_2 from "./ori2.jpg";
import ori_3 from "./ori3.jpg";
import ori_4 from "./ori4.jpg";
import ori_5 from "./ori5.jpg";

const imgs = [ori_1, ori_2, ori_3, ori_4, ori_5];

const About = () => {
  const [left, setLeft] = useState(0);

  useEffect(() => {
    const interval = setInterval(() => {
      setLeft((prevLeft) => (prevLeft - 1) % 5);
    }, 3000);

    return () => clearInterval(interval);
  }, []);

  return (
    <div
      style={{
        width: "100%",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <div
        style={{
          width: "500px",
          overflow: "hidden",
        }}
      >
        <div
          style={{
            width: "100%",
            display: "flex",
            justifyContent: "left",
            alignItems: "center",
            transform: `translateX(${left * 500}px)`,
            transition: "transform 0.5s ease-in-out",
          }}
        >
          {imgs.map((item) => {
            return (
              <img
                src={item}
                alt=""
                width="100%"
                height="100%"
                key={item}
                style={{
                  objectFit: "cover",
                  width: "500px",
                  height: "300px",
                }}
              />
            );
          })}
        </div>
      </div>
    </div>
  );
};

export default About;
