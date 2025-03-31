import styled from "@emotion/styled";

export { StyledTopRow, StyledTopRowFirstCell, StyledTopRowSecondCell, StyledTopRowThirdCell } from "./StyledTopRow/mod";
export {
  StyledMiddleRow,
  StyledMiddleRowFirstCell,
  StyledMiddleRowSecondCell,
  StyledMiddleRowThirdCell,
} from "./StyledMiddleRow/mod";
export {
  StyledBottomRow,
  StyledBottomRowFirstCell,
  StyledBottomRowSecondCell,
  StyledBottomRowThirdCell,
} from "./StyledBottomRow/mod";

export const StyledMatrixLayout = styled.section`
  display: grid;
  grid-template-columns: 1fr 1fr 1fr;
  grid-template-rows: 1fr 1fr 1fr;
  width: 300vw;
  height: 300vh;
  transform: ${(props) =>
    // @ts-ignore
    props.position ? `translate(-${props.position.x}, -${props.position.y})` : `translate(-100vw, -100vh)`};
`;
