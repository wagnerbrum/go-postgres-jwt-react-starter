import React from "react";
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";

import PrivateRoute from "./components/PrivateRoute";

import Navbar from "./components/Navbar";

import Login from "./pages/Login";
import Register from "./pages/Register";
import Home from "./pages/Home";
import CustumersList from "./pages/CustumersList";

const routeList = [
    { path: "/", exact: true, component: Login, protectedRoute: false },
    { path: "/login", exact: false, component: Login, protectedRoute: false },
    {
        path: "/register",
        exact: false,
        component: Register,
        protectedRoute: false,
    },
    { path: "/home", exact: false, component: Home, protectedRoute: true },
    { path: "/customers", exact: false, component: CustumersList, protectedRoute: true },
];

const Routes = () => (
    <Router>
        <Navbar />

        <Switch>{render}</Switch>
    </Router>
);

const render = routeList.map((item, key) =>
    item.protectedRoute ? (
        <PrivateRoute key={key} path={item.path} component={item.component} exact={item.exact} />
    ) : (
        <Route key={key} path={item.path} component={item.component} exact={item.exact} />
    )
);

export default Routes;
