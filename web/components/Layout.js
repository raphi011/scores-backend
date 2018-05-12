// @flow

import React from 'react';
import { withStyles } from 'material-ui/styles';

import Drawer from './Drawer';
import AppBar from '../containers/AppBarContainer';
import Snackbar from '../containers/SnackbarContainer';

import type { Player } from '../types';

type Props = {
  title: string,
  children: React$Element<>,
  userPlayer: Player,
  drawerOpen: boolean,
  onCloseDrawer: () => void,
  onOpenDrawer: () => void,
  onToggleGroup: () => void,
  classes: Object,
};

const styles = theme => ({
  style: {
    backgroundColor: theme.palette.background.paper,
    marginTop: '56px',
  },
});

const Layout = ({
  title,
  userPlayer,
  children,
  drawerOpen,
  onCloseDrawer,
  onOpenDrawer,
  onToggleGroup,
  groupOpen,
  classes,
}: Props) => (
  <div className={classes.style}>
    <Drawer
      userPlayer={userPlayer}
      onClose={onCloseDrawer}
      onOpen={onOpenDrawer}
      open={drawerOpen}
      onToggleGroup={onToggleGroup}
      groupOpen={groupOpen}
    />
    <AppBar onOpenMenu={onOpenDrawer} title={title} />
    {children}
    <Snackbar />
  </div>
);

export default withStyles(styles)(Layout);
