import React, { useState, useEffect } from "react";
import { Map, TileLayer, Marker } from "react-leaflet";
import { Icon } from "leaflet";
import moment from "moment";

import { FaEye, FaCcPaypal, FaCcAmex, FaCcVisa } from "react-icons/fa";
import { HiBadgeCheck } from "react-icons/hi";

import "leaflet/dist/leaflet.css";

import {
    Container,
    User,
    UserImage,
    UserImageTag,
    UserMaps,
    UserBody,
    BodyTitle,
    BodyTitleName,
    BodyTitleCity,
    BodyTag,
    BodyAddTag,
    ButtonAddTag,
    BodyPropertyContainer,
    BodyPropertyItemContainer,
    BodyPropertyItemLabel,
    BodyPropertyItemValue,
    Tables,
    TableCaption,
    TableTDStatusSuccess,
    TableTDStatusFail,
} from "./styles";

const initTimeStamp = 1512821643426;

const initListPayments = [
    {
        methodType: "CARD",
        coutry: "GBR",
        eWallet: "paypal",
        cardBin: "348444",
        lastFour: "9870",
        bank: "Natwest",
        registrationTime: initTimeStamp,
        successfulRegistration: true,
    },
    {
        methodType: "CARD",
        coutry: "GBR",
        eWallet: "amex",
        cardBin: "459005",
        lastFour: "2927",
        bank: "Barclays",
        registrationTime: initTimeStamp,
        successfulRegistration: false,
    },
    {
        methodType: "CARD",
        coutry: "GBR",
        eWallet: "visa",
        cardBin: "378854",
        lastFour: "9159",
        bank: "Barclays",
        registrationTime: initTimeStamp,
        successfulRegistration: false,
    },
];

const pointerIcon = new Icon({
    iconUrl: `${process.env.PUBLIC_URL}/assets/images/map-marker.svg`,
    iconRetinaUrl: `${process.env.PUBLIC_URL}/assets/images/map-marker.svg`,
    iconAnchor: [15, 35],
    iconSize: [25, 55],
});

const CustumersList = () => {
    const PER_PAGE = 8;

    const [data, setData] = useState([]);
    const [dataPagesCount, setDataPagesCount] = useState(0);
    const [dataFiltered, setDataFiltered] = useState([]);
    const [currentPage, setCurrentPage] = useState(0);

    useEffect(() => {
        handlePopulateList();
    }, []);

    useEffect(() => {
        handleFilterList();
        handlePageCount();
    }, [data, currentPage]);

    const handleRandomInt = (min, max) => {
        min = Math.ceil(min);
        max = Math.floor(max);
        return Math.floor(Math.random() * (max - min)) + min;
    };

    const handleGenerateLargeList = (list, length) => {
        const newList = [];

        for (var i = 0; i < length; i++) {
            for (var j = 0; j < list.length; j++) {
                const data = {
                    ...list[j],
                    registrationTime: list[j].registrationTime + 100000 * handleRandomInt(1, 1000),
                };

                newList.push(data);
            }
        }

        return newList;
    };

    const handleFilterList = () => {
        const offset = currentPage * PER_PAGE;

        const currentPageData = data.slice(offset, offset + PER_PAGE);

        setDataFiltered(currentPageData);
    };

    const handleOrderData = (data) => {
        return data.sort((a, b) => b.registrationTime - a.registrationTime);
    };

    const handlePopulateList = () => {
        const generatedList = handleGenerateLargeList(initListPayments, 10);
        const ordenatedList = handleOrderData(generatedList);

        setData(ordenatedList);
    };

    const handlePageCount = () => {
        setDataPagesCount(Math.ceil(data.length / PER_PAGE));
    };

    const handleNextPage = () => {
        if (currentPage + 1 < dataPagesCount) {
            setCurrentPage(currentPage + 1);
        }
    };

    const handlePreviousPage = () => {
        if (currentPage + 1 > 1) {
            setCurrentPage(currentPage - 1);
        }
    };

    const positon = [-29.9153717, -51.0645172];

    const handleCardIcon = (eWallet) => {
        switch (eWallet) {
            case "paypal": {
                return <FaCcPaypal color={"#2b74b6"} size={24} className="mr-2" />;
            }
            case "amex": {
                return <FaCcAmex color={"#2b74b6"} size={24} className="mr-2" />;
            }
            case "visa": {
                return <FaCcVisa color={"#2b74b6"} size={24} className="mr-2" />;
            }
            default: {
                return <FaCcVisa color={"#2b74b6"} size={24} className="mr-2" />;
            }
        }
    };

    return (
        <Container className="container-fluid">
            <div className="row px-4">
                <User className="mx-auto">
                    <UserImage>
                        ZM
                        <UserImageTag>100</UserImageTag>
                    </UserImage>

                    <UserMaps>
                        <Map center={positon} zoom={15} style={{ width: "100%", height: "100%" }} zoomControl={false}>
                            <TileLayer
                                url={`https://api.mapbox.com/styles/v1/mapbox/light-v10/tiles/256/{z}/{x}/{y}@2x?access_token=${process.env.REACT_APP_MAPBOX_TOKEN}`}
                            />

                            <Marker position={positon} icon={pointerIcon}></Marker>
                        </Map>
                    </UserMaps>

                    <UserBody>
                        <BodyTitle className="row">
                            <div className="col-9">
                                <BodyTitleName>Zachar Morgany</BodyTitleName>
                                <BodyTitleCity>CustZachary</BodyTitleCity>
                            </div>
                            <div className="col-3 pr-4 text-right">
                                <button className="btn btn-outline-secondary pt-2 mt-2">
                                    <FaEye />
                                </button>
                            </div>
                        </BodyTitle>

                        <div className="my-3"></div>

                        <BodyTag className="row">
                            d+zachm@rexample.com
                            <HiBadgeCheck className="ml-2" color={"#99c7f4"} />
                        </BodyTag>

                        <BodyTag className="row">
                            +447874125478
                            <HiBadgeCheck className="ml-2" color={"#99c7f4"} />
                        </BodyTag>

                        <div className="my-3"></div>

                        <BodyAddTag>
                            <ButtonAddTag>+ add tag</ButtonAddTag>
                        </BodyAddTag>

                        <div className="my-3"></div>

                        <BodyPropertyContainer className="row">
                            <BodyPropertyItemContainer className="col-6 mb-3">
                                <BodyPropertyItemLabel>Account age</BodyPropertyItemLabel>
                                <BodyPropertyItemValue>1 day</BodyPropertyItemValue>
                            </BodyPropertyItemContainer>

                            <BodyPropertyItemContainer className="col-6 mb-3">
                                <BodyPropertyItemLabel>Average order value</BodyPropertyItemLabel>
                                <BodyPropertyItemValue>247</BodyPropertyItemValue>
                            </BodyPropertyItemContainer>

                            <BodyPropertyItemContainer className="col-6 mb-3">
                                <BodyPropertyItemLabel>Total lifetime spend</BodyPropertyItemLabel>
                                <BodyPropertyItemValue>494</BodyPropertyItemValue>
                            </BodyPropertyItemContainer>
                        </BodyPropertyContainer>
                    </UserBody>
                </User>
                <Tables className="p-3 mt-4">
                    <div className="table-responsive">
                        <table className="table table-sm table-striped mb-2">
                            <thead>
                                <tr>
                                    <th className="pl-3" scope="col">
                                        TYPE
                                    </th>
                                    <th scope="col">DETAILS</th>
                                    <th scope="col">COUNTRY</th>
                                    <th scope="col">STATUS</th>
                                    <th className="pr-3" scope="col">
                                        REGISTRATION TIME
                                    </th>
                                </tr>
                            </thead>
                            <tbody>
                                {dataFiltered.map((value, key) => (
                                    <tr key={key}>
                                        <td className="pl-3">{value.methodType}</td>
                                        <td className="d-flex">
                                            {handleCardIcon(value.eWallet)}
                                            {`${value.cardBin.substr(0, 4)} ${value.cardBin.substr(4, 6)}•• •••• ${
                                                value.lastFour
                                            } (${value.bank})`}
                                        </td>
                                        <td>{value.coutry}</td>
                                        <td>
                                            {value.successfulRegistration ? (
                                                <TableTDStatusSuccess>Success</TableTDStatusSuccess>
                                            ) : (
                                                <TableTDStatusFail>Failed</TableTDStatusFail>
                                            )}
                                        </td>
                                        <td className="pr-3">
                                            {moment(value.registrationTime).format("DD MMM, YYYY, h:mm a")}
                                        </td>
                                    </tr>
                                ))}
                            </tbody>
                        </table>
                        <TableCaption>
                            <div className="cursor-pointer" onClick={() => handlePreviousPage()}>
                                Prev
                            </div>
                            <div>
                                {currentPage + 1}/{dataPagesCount}
                            </div>
                            <div className="cursor-pointer" onClick={() => handleNextPage()}>
                                Next
                            </div>
                        </TableCaption>
                    </div>
                </Tables>
            </div>
        </Container>
    );
};

export default CustumersList;
