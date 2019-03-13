using FileBackup.ViewModels;
using FileBackup.Views;
using System.Windows;

namespace FileBackup
{
    /// <summary>
    /// Interaction logic for App.xaml
    /// </summary>
    public partial class App : Application
    {
        private System.Windows.Forms.NotifyIcon _NotifyIcon;
        MainView _MainView;
        MainViewModel _MainViewModel;

        /// <summary>
        /// 
        /// </summary>
        /// <param name="e"></param>
        protected override void OnStartup(StartupEventArgs e)
        {
            base.OnStartup(e);

            //
            _MainView = new MainView();
            _MainViewModel = new MainViewModel();
            _MainView.DataContext = _MainViewModel;

            //
            _NotifyIcon = new System.Windows.Forms.NotifyIcon();
            _NotifyIcon.Icon = FileBackup.Properties.Resources.NotifyIconImage;
            _NotifyIcon.Visible = true;

            _NotifyIcon.MouseDoubleClick += NotifyIconMouseDoubleClick;
        }

        private void NotifyIconMouseDoubleClick(object sender, System.Windows.Forms.MouseEventArgs e)
        {
            _MainView.ShowDialog();
        }
    }
}
