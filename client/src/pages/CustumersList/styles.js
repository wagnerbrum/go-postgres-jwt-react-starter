import styled from "styled-components";

export const Container = styled.div`
    background-color: #eee;

    height: calc(100vh - 50px);
`;

export const User = styled.div`
    min-height: 600px;
    max-width: 450px;

    display: flex;
    flex-direction: column;

    position: relative;

    background-color: #fff;
`;

export const UserImage = styled.div`
    height: 80px;
    width: 80px;

    border: 4px solid #fff;
    border-radius: 40px;
    background-color: #d5e4bf;
    color: #fff;
    font-size: 2em;
    font-weight: 600;

    display: flex;
    justify-content: center;
    align-items: center;

    position: absolute;
    left: 30px;
    top: 140px;

    z-index: 10;
`;

export const UserImageTag = styled.div`
    height: 30px;
    padding-left: 4px;
    padding-right: 4px;
    padding-bottom: 4px;

    border-radius: 1px;
    background-color: #da4c62;
    color: #fff;
    font-size: 0.9em;
    font-weight: 500;

    display: flex;
    justify-content: center;
    align-items: center;

    position: absolute;
    left: 65px;
    top: 23px;

    z-index: 11;
`;

export const UserMaps = styled.div`
    height: 200px;
    position: relative;
    z-index: 5;
`;

export const UserBody = styled.div`
    padding-top: 40px;
    padding-right: 15px;
    padding-bottom: 40px;
    padding-left: 15px;
`;

export const BodyTitle = styled.div``;

export const BodyTitleName = styled.div`
    font-size: 1.5em;
    font-weight: 500;
`;

export const BodyTitleCity = styled.div`
    font-size: 1.1em;
    color: #ccc;
`;

export const BodyTag = styled.div`
    display: flex;
    justify-content: start;
    align-items: center;

    padding-right: 15px;
    padding-left: 15px;

    font-weight: 400;
    font-size: 1em;
    letter-spacing: 1px;
    color: #333;
`;

export const BodyAddTag = styled.div`
    display: flex;
`;

export const ButtonAddTag = styled.div`
    color: #aaa;
    font-weight: 500;
    background-color: #ddd;

    margin-right: 15px;
    padding: 2px 7px 2px 7px;
    border-radius: 5px;

    cursor: pointer;
`;

export const BodyPropertyContainer = styled.div``;

export const BodyPropertyItemContainer = styled.div``;

export const BodyPropertyItemLabel = styled.div`
    color: #aaabbb;
`;

export const BodyPropertyItemValue = styled.div``;

export const Tables = styled.div`
    display: flex;
    flex: 1;
    background-color: #fff;
    margin-left: 20px;

    & table {
        & tbody {
            font-size: 0.9em;
            letter-spacing: 1px;
        }
    }
`;

export const TableCaption = styled.div`
    display: flex;
    flex: 1;
    justify-content: space-between;
    padding-left: 30px;
    padding-right: 30px;
    color: #aaa;
`;

export const TableTDStatus = styled.div`
    display: inline-block;

    border-radius: 5px;
    padding: 2px 5px 2px 5px;
`;

export const TableTDStatusSuccess = styled(TableTDStatus)`
    background-color: #d4edda;
    color: #155724;
`;

export const TableTDStatusFail = styled(TableTDStatus)`
    color: #721c24;
    background-color: #f8d7da;
`;
