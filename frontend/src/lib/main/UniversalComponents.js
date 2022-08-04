import AutoComplete from "primevue/autocomplete";
import Accordion from "primevue/accordion";
import AccordionTab from "primevue/accordiontab";
import Avatar from "primevue/avatar";
import AvatarGroup from "primevue/avatargroup";
import Badge from "primevue/badge";
import Button from "primevue/button";
import Breadcrumb from "primevue/breadcrumb";
import Calendar from "primevue/calendar";
import Card from "primevue/card";
import Carousel from "primevue/carousel";
import Chart from "primevue/chart";
import Checkbox from "primevue/checkbox";
import Chip from "primevue/chip";
import Chips from "primevue/chips";
import ColorPicker from "primevue/colorpicker";
import Column from "primevue/column";
import ConfirmDialog from "primevue/confirmdialog";
import ConfirmPopup from "primevue/confirmpopup";
import ContextMenu from "primevue/contextmenu";
import DataTable from "primevue/datatable";
import DataView from "primevue/dataview";
import DataViewLayoutOptions from "primevue/dataviewlayoutoptions";
import Dialog from "primevue/dialog";
import Divider from "primevue/divider";
import Dropdown from "primevue/dropdown";
import Fieldset from "primevue/fieldset";
import FileUpload from "primevue/fileupload";
import Image from "primevue/image";
import InlineMessage from "primevue/inlinemessage";
import Inplace from "primevue/inplace";
import InputMask from "primevue/inputmask";
import InputNumber from "primevue/inputnumber";
import InputSwitch from "primevue/inputswitch";
import InputText from "primevue/inputtext";
import Knob from "primevue/knob";
import Galleria from "primevue/galleria";
import Listbox from "primevue/listbox";
import MegaMenu from "primevue/megamenu";
import Menu from "primevue/menu";
import Menubar from "primevue/menubar";
import Message from "primevue/message";
import MultiSelect from "primevue/multiselect";
import OrderList from "primevue/orderlist";
import OrganizationChart from "primevue/organizationchart";
import OverlayPanel from "primevue/overlaypanel";
import Paginator from "primevue/paginator";
import Panel from "primevue/panel";
import PanelMenu from "primevue/panelmenu";
import Password from "primevue/password";
import PickList from "primevue/picklist";
import ProgressBar from "primevue/progressbar";
import Rating from "primevue/rating";
import RadioButton from "primevue/radiobutton";
import SelectButton from "primevue/selectbutton";
import ScrollPanel from "primevue/scrollpanel";
import ScrollTop from "primevue/scrolltop";
import Slider from "primevue/slider";
import Sidebar from "primevue/sidebar";
import Skeleton from "primevue/skeleton";
import SplitButton from "primevue/splitbutton";
import Splitter from "primevue/splitter";
import SplitterPanel from "primevue/splitterpanel";
import Steps from "primevue/steps";
import TabMenu from "primevue/tabmenu";
import Tag from "primevue/tag";
import TieredMenu from "primevue/tieredmenu";
import Textarea from "primevue/textarea";
import Timeline from "primevue/timeline";
import Toast from "primevue/toast";
import Toolbar from "primevue/toolbar";
import TabView from "primevue/tabview";
import TabPanel from "primevue/tabpanel";
import ToggleButton from "primevue/togglebutton";
import TreeSelect from "primevue/treeselect";
import TreeTable from "primevue/treetable";
import Tree from "primevue/tree";
import TriStateCheckbox from "primevue/tristatecheckbox";

import BlockViewer from "../../components/app/BlockViewer";
import AppInputStyleSwitch from "../../components/app/AppInputStyleSwitch";
import FormPage from "../../components/common/FormPage.vue";
import InvalidPage from "../../components/common/InvalidPage.vue";

// you cannot just use dynamic imports here
// if you do, webpack will create a small file for each import
// it will take much longer to load all of them
const UniversalComponents = {
  Accordion: Accordion,
  AccordionTab: AccordionTab,
  AutoComplete: AutoComplete,
  Avatar: Avatar,
  AvatarGroup: AvatarGroup,
  Badge: Badge,
  Breadcrumb: Breadcrumb,
  Button: Button,
  Calendar: Calendar,
  Card: Card,
  Carousel: Carousel,
  Chart: Chart,
  Checkbox: Checkbox,
  Chip: Chip,
  Chips: Chips,
  ColorPicker: ColorPicker,
  Column: Column,
  ConfirmDialog: ConfirmDialog,
  ConfirmPopup: ConfirmPopup,
  ContextMenu: ContextMenu,
  DataTable: DataTable,
  DataView: DataView,
  DataViewLayoutOptions: DataViewLayoutOptions,
  Dialog: Dialog,
  Divider: Divider,
  Dropdown: Dropdown,
  Fieldset: Fieldset,
  FileUpload: FileUpload,
  Image: Image,
  InlineMessage: InlineMessage,
  Inplace: Inplace,
  InputMask: InputMask,
  InputNumber: InputNumber,
  InputSwitch: InputSwitch,
  InputText: InputText,
  Galleria: Galleria,
  Knob: Knob,
  Listbox: Listbox,
  MegaMenu: MegaMenu,
  Menu: Menu,
  Menubar: Menubar,
  Message: Message,
  MultiSelect: MultiSelect,
  OrderList: OrderList,
  OrganizationChart: OrganizationChart,
  OverlayPanel: OverlayPanel,
  Paginator: Paginator,
  Panel: Panel,
  PanelMenu: PanelMenu,
  Password: Password,
  PickList: PickList,
  ProgressBar: ProgressBar,
  RadioButton: RadioButton,
  Rating: Rating,
  SelectButton: SelectButton,
  ScrollPanel: ScrollPanel,
  ScrollTop: ScrollTop,
  Slider: Slider,
  Sidebar: Sidebar,
  Skeleton: Skeleton,
  SplitButton: SplitButton,
  Splitter: Splitter,
  SplitterPanel: SplitterPanel,
  Steps: Steps,
  TabMenu: TabMenu,
  TabView: TabView,
  TabPanel: TabPanel,
  Tag: Tag,
  Textarea: Textarea,
  TieredMenu: TieredMenu,
  Timeline: Timeline,
  Toast: Toast,
  Toolbar: Toolbar,
  ToggleButton: ToggleButton,
  Tree: Tree,
  TreeSelect: TreeSelect,
  TreeTable: TreeTable,
  TriStateCheckbox: TriStateCheckbox,
  //
  // custom components
  //
  AppInputStyleSwitch: AppInputStyleSwitch,
  BlockViewer: BlockViewer,
  FormPage: FormPage,
  InvalidPage: InvalidPage,
};

export default UniversalComponents;
