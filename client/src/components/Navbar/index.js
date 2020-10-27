import React from "react";
import { useLocation, Link, useHistory } from "react-router-dom";

import { isAuth, deleteToken } from "../../shared/utils";

import { Container, List, Item, Logout } from "./styles";

const listNavbarItens = [
    { label: "Home", path: "/home" },
    { label: "Customers", path: "/customers" },
];

const Navbar = () => {
    const { pathname } = useLocation();
    const history = useHistory();

    const handleItemSelected = (path) => {
        const pathnameSplit = pathname.split("/");

        if (pathname === path) {
            return true;
        }

        if (pathnameSplit.length > 1) {
            if (`/${pathnameSplit[1]}` === path) {
                return true;
            }
        }

        return false;
    };

    const handleLogout = () => {
        deleteToken();
        history.push("/");
    };

    return (
        isAuth() && (
            <Container className="container-fluid">
                <List>
                    {listNavbarItens.map((item, key) => (
                        <Item key={key} selected={handleItemSelected(item.path)}>
                            <Link to={item.path}>{item.label}</Link>
                        </Item>
                    ))}
                </List>
                <Logout onClick={() => handleLogout()}>Sair</Logout>
            </Container>
        )
    );
};

export default Navbar;
