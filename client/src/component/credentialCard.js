import Button from "react-bootstrap/Button";
import Card from "react-bootstrap/Card";
import { Row, Col } from "react-bootstrap";
import { useState } from "react";

const cardStyle = {
  background: "linear-gradient(to bottom left, #0000ff 0%, #003300 88%)",
  // background:"#ff6b9f",
  width: "20rem",
  justifyContent: "center",
  border: "solid green 1px",
  //  backgroundColor:"#e6fff7"
  color: "white",
};

function CredentialCard({ name, uid, cid }) {
  // console.log(id)
  const [unlockButton, setUnlockButton] = useState(false);
  const [unlock, setUnlock] = useState(false);

  return (
    <Card style={cardStyle} key={uid}>
      <Card.Body flex flexDirection="row">
        <Row>
          <Col style={{ color: "white" }}>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="26"
              height="24"
              fill="currentColor"
              class="bi bi-wallet2"
              viewBox="0 0 18 18"
            >
              <path d="M12.136.326A1.5 1.5 0 0 1 14 1.78V3h.5A1.5 1.5 0 0 1 16 4.5v9a1.5 1.5 0 0 1-1.5 1.5h-13A1.5 1.5 0 0 1 0 13.5v-9a1.5 1.5 0 0 1 1.432-1.499zM5.562 3H13V1.78a.5.5 0 0 0-.621-.484zM1.5 4a.5.5 0 0 0-.5.5v9a.5.5 0 0 0 .5.5h13a.5.5 0 0 0 .5-.5v-9a.5.5 0 0 0-.5-.5z" />
            </svg>
          </Col>
          <Col>
            <Card.Title>{name}</Card.Title>
          </Col>
          <Col></Col>
        </Row>
        <Button
          style={{ border: "solid pink 1px" }}
          variant="success"
          size="sm"
          href={`https://orange-informal-leopon-463.mypinata.cloud/ipfs/${cid}`}
          target="blank"
        >
          Check Credential
        </Button>
      </Card.Body>
    </Card>
  );
}

export default CredentialCard;
