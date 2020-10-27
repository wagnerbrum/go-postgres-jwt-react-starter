import styled from "styled-components";

export const Container = styled.div`
    background-color: #98a9bd;
    color: #fff;

    display: flex;
    height: 50px;
    justify-content: space-between;
    padding-left: 30px;
    padding-right: 30px;
    align-items: center;
`;

export const List = styled.div`
    display: flex;
    flex-direction: row;

    padding-top: 15px;
    padding-bottom: 15px;
`;

export const Item = styled.div`
    padding-top: 5px;
    padding-bottom: 5px;
    padding-left: 10px;
    padding-right: 10px;
    margin-right: 10px;

    font-weight: 600;

    background-color: ${(props) => (props.selected ? "#8494a5" : "transparent")};
    border-radius: 5px;

    &:hover {
        background-color: #8494a5;
    }

    a {
        color: #fff;

        &:hover {
            text-decoration: none;
        }
    }

    cursor: pointer;
`;

export const Logout = styled.div`
    color: #fff;
    font-weight: bold;
    letter-spacing: 2px;
    cursor: pointer;
`;
